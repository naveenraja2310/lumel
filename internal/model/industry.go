package model

import (
	"time"
)

type Industry struct {
	ID          string    `json:"id" bson:"_id,omitempty"`
	Name        string    `json:"name" bson:"name"`
	Description string    `json:"description" bson:"description"`
	Image       string    `json:"image" bson:"image"`
	Slug        string    `json:"slug" bson:"slug"`
	CreatedAt   time.Time `json:"createdat" bson:"createdat"`
	UpdatedAt   time.Time `json:"updatedat" bson:"updatedat"`
}

func (i Industry) TableName() string {
	return "industry"
}
