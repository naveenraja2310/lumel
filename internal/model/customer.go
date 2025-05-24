package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Customer struct {
	ID         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	CustomerID string             `json:"customer_id" bson:"customer_id" validate:"required"`
	Name       string             `json:"name" bson:"name" validate:"required"`
	Email      string             `json:"email" bson:"email" validate:"required,email"`
	Address    string             `json:"address" bson:"address"`
	CreateAt   time.Time          `json:"created_at" bson:"created_at"`
	UpdateAt   time.Time          `json:"updated_at" bson:"updated_at"`
}

func (Customer) TableName() string {
	return "customers"
}
