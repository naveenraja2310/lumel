package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID            primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	OrderID       string             `json:"order_id" bson:"order_id" validate:"required"`
	CustomerID    primitive.ObjectID `json:"customer_id" bson:"customer_id" validate:"required"` // reference to customer
	ProductID     primitive.ObjectID `json:"product_id" bson:"product_id" validate:"required"`   // reference to product
	Region        string             `json:"region" bson:"region" validate:"required"`
	DateOfSale    time.Time          `json:"date_of_sale" bson:"date_of_sale" validate:"required"`
	QuantitySold  int                `json:"quantity_sold" bson:"quantity_sold" validate:"required,min=1"`
	UnitPrice     float64            `json:"unit_price" bson:"unit_price" validate:"required,min=0"`
	Discount      float64            `json:"discount" bson:"discount" validate:"min=0,max=1"`
	ShippingCost  float64            `json:"shipping_cost" bson:"shipping_cost" validate:"min=0"`
	PaymentMethod string             `json:"payment_method" bson:"payment_method" validate:"required"`
	TotalAmount   float64            `json:"total_amount" bson:"total_amount"`
	CreateAt      time.Time          `json:"created_at" bson:"created_at"`
	UpdateAt      time.Time          `json:"updated_at" bson:"updated_at"`
}

// OrderWithDetails represents order with populated customer and product details
type OrderWithDetails struct {
	ID            string    `json:"id" bson:"_id,omitempty"`
	OrderID       string    `json:"order_id" bson:"order_id"`
	Customer      Customer  `json:"customer" bson:"customer"`
	Product       Product   `json:"product" bson:"product"`
	Region        string    `json:"region" bson:"region"`
	DateOfSale    time.Time `json:"date_of_sale" bson:"date_of_sale"`
	QuantitySold  int       `json:"quantity_sold" bson:"quantity_sold"`
	UnitPrice     float64   `json:"unit_price" bson:"unit_price"`
	Discount      float64   `json:"discount" bson:"discount"`
	ShippingCost  float64   `json:"shipping_cost" bson:"shipping_cost"`
	PaymentMethod string    `json:"payment_method" bson:"payment_method"`
	TotalAmount   float64   `json:"total_amount" bson:"total_amount"`
	CreateAt      time.Time `json:"created_at" bson:"created_at"`
	UpdateAt      time.Time `json:"updated_at" bson:"updated_at"`
}

func (Order) TableName() string {
	return "orders"
}
