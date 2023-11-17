package model

import (
	"common"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MenuItem struct {
	ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	RestaurantID int                `json:"restaurant_id" bson:"restaurant_id" validate:"required"`
	Name         string             `json:"name" bson:"name" validate:"required"`
	Description  string             `json:"description" bson:"description"`
	Image        *common.Image      `json:"image" bson:"image" validate:"required"`
	Price        float32            `json:"price" bson:"price" validate:"required"`
	Addon        *[]AddOn           `json:"add_on" bson:"add_on,omitempty"`
	UpdatedAt    time.Time          `json:"updated_at,omitempty" bson:"updated_at"`
	CreatedAt    time.Time          `json:"created_at,omitempty" bson:"created_at"`
}

// Implement bson.Marshaler, and MarshalBSON() function will be called when you save values of *MenuItem type.
func (m *MenuItem) MarshalBSON() ([]byte, error) {
	if m.CreatedAt.IsZero() {
		m.CreatedAt = time.Now()
	}
	m.UpdatedAt = time.Now()

	type my MenuItem
	return bson.Marshal((*my)(m))
}

type AddOn struct {
	Name  string        `json:"name" bson:"name" validate:"required"`
	Image *common.Image `json:"image" bson:"image" validate:"required"`
	Price float32       `json:"price" bson:"price" validate:"required"`
}
