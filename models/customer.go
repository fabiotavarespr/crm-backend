package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CustomerCreateRequest struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Name      string             `json:"name" bson:"name" binding:"required"`
	Email     string             `json:"email" bson:"email" binding:"required"`
	Role      string             `json:"role" bson:"role"`
	Phone     string             `json:"phone" bson:"phone"`
	Contacted bool               `json:"contacted" bson:"contacted"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

type CustomerCreateResponse struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Name      string             `json:"name" bson:"name" binding:"required"`
	Email     string             `json:"email" bson:"email" binding:"required"`
	Role      string             `json:"role" bson:"role"`
	Phone     string             `json:"phone" bson:"phone"`
	Contacted bool               `json:"contacted" bson:"contacted"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

type CustomerGetResponse struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Name      string             `json:"name" bson:"name" binding:"required"`
	Email     string             `json:"email" bson:"email" binding:"required"`
	Role      string             `json:"role" bson:"role"`
	Phone     string             `json:"phone" bson:"phone"`
	Contacted bool               `json:"contacted" bson:"contacted"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

type CustomersGetResponse struct {
	Customers []CustomerGetResponse `json:"customers" bson:"customers" `
}

type CustomerUpdateRequest struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Name      string             `json:"name" bson:"name" binding:"required"`
	Email     string             `json:"email" bson:"email" binding:"required"`
	Role      string             `json:"role" bson:"role"`
	Phone     string             `json:"phone" bson:"phone"`
	Contacted bool               `json:"contacted" bson:"contacted"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

type CustomerUpdateResponse struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Name      string             `json:"name" bson:"name" binding:"required"`
	Email     string             `json:"email" bson:"email" binding:"required"`
	Role      string             `json:"role" bson:"role"`
	Phone     string             `json:"phone" bson:"phone"`
	Contacted bool               `json:"contacted" bson:"contacted"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}
