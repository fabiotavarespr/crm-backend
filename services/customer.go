package services

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/fabiotavarespr/crm-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CustomerService interface {
	AddCustomer(*models.CustomerCreateRequest) (*models.CustomerCreateResponse, error)
	GetCustomers() (*models.CustomersGetResponse, error)
	GetCustomer(string) (*models.CustomerGetResponse, error)
	UpdateCustomer(string, *models.CustomerUpdateRequest) (*models.CustomerUpdateResponse, error)
	DeleteCustomer(string) error
}

type CustomerServiceImpl struct {
	collection *mongo.Collection
	ctx        context.Context
}

func NewCustomerService(collection *mongo.Collection, ctx context.Context) CustomerService {
	return &CustomerServiceImpl{collection, ctx}
}

func (cs *CustomerServiceImpl) AddCustomer(customer *models.CustomerCreateRequest) (*models.CustomerCreateResponse, error) {
	customer.ID = primitive.NewObjectID()
	customer.CreatedAt = time.Now()
	customer.UpdatedAt = customer.CreatedAt
	customer.Email = strings.ToLower(customer.Email)

	res, err := cs.collection.InsertOne(cs.ctx, &customer)

	if err != nil {
		if er, ok := err.(mongo.WriteException); ok && er.WriteErrors[0].Code == 11000 {
			return nil, errors.New("Customer with that email already exist")
		}
		return nil, err
	}

	// Create a unique index for the email field
	opt := options.Index()
	opt.SetUnique(true)
	index := mongo.IndexModel{Keys: bson.M{"email": 1}, Options: opt}

	if _, err := cs.collection.Indexes().CreateOne(cs.ctx, index); err != nil {
		return nil, errors.New("Could not create index for email")
	}

	var newCustomer *models.CustomerCreateResponse
	query := bson.M{"_id": res.InsertedID}

	err = cs.collection.FindOne(cs.ctx, query).Decode(&newCustomer)
	if err != nil {
		return nil, err
	}

	return newCustomer, nil
}

func (cs *CustomerServiceImpl) GetCustomers() (*models.CustomersGetResponse, error) {
	newCustomers := make([]models.CustomerGetResponse, 0)
	res, err := cs.collection.Find(cs.ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	defer res.Close(cs.ctx)

	for res.Next(cs.ctx) {
		var customer models.CustomerGetResponse
		if err = res.Decode(&customer); err != nil {
			return nil, err
		}
		newCustomers = append(newCustomers, customer)
	}

	return &models.CustomersGetResponse{
		Customers: newCustomers,
	}, nil
}

func (cs *CustomerServiceImpl) GetCustomer(id string) (*models.CustomerGetResponse, error) {
	var newCustomer *models.CustomerGetResponse

	objectID, _ := primitive.ObjectIDFromHex(id)
	query := bson.M{"_id": objectID}

	err := cs.collection.FindOne(cs.ctx, query).Decode(&newCustomer)
	if err != nil {
		return nil, err
	}

	return newCustomer, nil

}

func (cs *CustomerServiceImpl) DeleteCustomer(id string) error {

	objectID, _ := primitive.ObjectIDFromHex(id)
	query := bson.M{"_id": objectID}

	deletedCustomer := cs.collection.FindOneAndDelete(cs.ctx, query)
	if deletedCustomer.Err() != nil {
		return deletedCustomer.Err()
	}

	return nil
}

func (cs *CustomerServiceImpl) UpdateCustomer(id string, customer *models.CustomerUpdateRequest) (*models.CustomerUpdateResponse, error) {
	objectID, _ := primitive.ObjectIDFromHex(id)
	query := bson.M{"_id": objectID}

	customer.UpdatedAt = time.Now()
	customer.ID = objectID

	update := bson.M{
		"$set": customer,
	}

	result := cs.collection.FindOneAndUpdate(cs.ctx, query, update)
	if result.Err() != nil {
		if strings.Contains(result.Err().Error(), "E11000") {
			return nil, errors.New("Customer with that email already exist")
		}
		return nil, result.Err()
	}

	var updatedCustomer *models.CustomerUpdateResponse

	err := cs.collection.FindOne(cs.ctx, query).Decode(&updatedCustomer)
	if err != nil {
		return nil, err
	}

	return updatedCustomer, nil
}
