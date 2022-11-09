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
	AddCustomer(*models.CustomerRequest) (*models.CustomerResponse, error)
}

type CustomerServiceImpl struct {
	collection *mongo.Collection
	ctx        context.Context
}

func NewCustomerService(collection *mongo.Collection, ctx context.Context) CustomerService {
	return &CustomerServiceImpl{collection, ctx}
}

func (cs *CustomerServiceImpl) AddCustomer(customer *models.CustomerRequest) (*models.CustomerResponse, error) {
	customer.ID = primitive.NewObjectID()
	customer.CreatedAt = time.Now()
	customer.UpdatedAt = customer.CreatedAt
	customer.Email = strings.ToLower(customer.Email)

	res, err := cs.collection.InsertOne(cs.ctx, &customer)

	if err != nil {
		if er, ok := err.(mongo.WriteException); ok && er.WriteErrors[0].Code == 11000 {
			return nil, errors.New("customer with that email already exist")
		}
		return nil, err
	}

	// Create a unique index for the email field
	opt := options.Index()
	opt.SetUnique(true)
	index := mongo.IndexModel{Keys: bson.M{"email": 1}, Options: opt}

	if _, err := cs.collection.Indexes().CreateOne(cs.ctx, index); err != nil {
		return nil, errors.New("could not create index for email")
	}

	var newCustomer *models.CustomerResponse
	query := bson.M{"_id": res.InsertedID}

	err = cs.collection.FindOne(cs.ctx, query).Decode(&newCustomer)
	if err != nil {
		return nil, err
	}

	return newCustomer, nil
}