package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CustomerRequest struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Name      string             `json:"name" bson:"name" binding:"required"`
	Email     string             `json:"email" bson:"email" binding:"required"`
	Role      string             `json:"role" bson:"role"`
	Phone     string             `json:"phone" bson:"phone"`
	Contacted bool               `json:"contacted" bson:"contacted"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

type CustomerResponse struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Name      string             `json:"name" bson:"name" binding:"required"`
	Email     string             `json:"email" bson:"email" binding:"required"`
	Role      string             `json:"role" bson:"role"`
	Phone     string             `json:"phone" bson:"phone"`
	Contacted bool               `json:"contacted" bson:"contacted"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}