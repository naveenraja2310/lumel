package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	ProductID   string             `json:"product_id" bson:"product_id" validate:"required"`
	Name        string             `json:"name" bson:"name" validate:"required"`
	Category    string             `json:"category" bson:"category" validate:"required"`
	Description string             `json:"description" bson:"description"`
	CreateAt    time.Time          `json:"created_at" bson:"created_at"`
	UpdateAt    time.Time          `json:"updated_at" bson:"updated_at"`
}

func (Product) TableName() string {
	return "products"
}
