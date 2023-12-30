package model

import (
	"common"
	"errors"
	"time"
)

// type MenuItem struct {
// 	ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
// 	RestaurantID int                `json:"restaurant_id" bson:"restaurant_id" validate:"required"`
// 	Name         string             `json:"name" bson:"name" validate:"required"`
// 	Description  string             `json:"description" bson:"description"`
// 	Image        *common.Image      `json:"image" bson:"image" validate:"required"`
// 	Price        float32            `json:"price" bson:"price" validate:"required"`
// 	Category     string             `json:"category" bson:"category" validate:"required"`
// 	UpdatedAt    time.Time          `json:"updated_at,omitempty" bson:"updated_at"`
// 	CreatedAt    time.Time          `json:"created_at,omitempty" bson:"created_at"`
// }

type MenuItem struct {
	common.SqlModel
	RestaurantID int           `json:"restaurant_id" gorm:"not null"`
	Name         string        `json:"name" gorm:"not null"`
	Description  string        `json:"description"`
	Image        *common.Image `json:"image" gorm:"not null"`
	Price        float32       `json:"price" gorm:"not null"`
	Category     string        `json:"category" gorm:"not null"`
}

// Implement bson.Marshaler, and MarshalBSON() function will be called when you save values of *MenuItem type.
// func (m *MenuItem) MarshalBSON() ([]byte, error) {
// 	if m.CreatedAt.IsZero() {
// 		m.CreatedAt = time.Now()
// 	}
// 	m.UpdatedAt = time.Now()

// 	type my MenuItem
// 	return bson.Marshal((*my)(m))
// }

//	type UpdateMenuItem struct {
//		Name        string        `json:"name" bson:"name,omitempty"`
//		Description *string       `json:"description" bson:"description,omitempty"`
//		Image       *common.Image `json:"image" bson:"image,omitempty"`
//		Price       *float32      `json:"price" bson:"price,omitempty"`
//		UpdatedAt   time.Time     `json:"updated_at,omitempty" bson:"updated_at"`
//	}
type UpdateMenuItem struct {
	Name        string        `json:"name" gorm:"column:name"`
	Description *string       `json:"description" gorm:"column:description"`
	Image       *common.Image `json:"image" gorm:"column:image"`
	Price       *float32      `json:"price" gorm:"column:price"`
	UpdatedAt   time.Time     `json:"updated_at" gorm:"column:updated_at"`
}

// func (m *UpdateMenuItem) MarshalBSON() ([]byte, error) {
// 	m.UpdatedAt = time.Now()

// 	type my UpdateMenuItem
// 	return bson.Marshal((*my)(m))
// }

const (
	CategoryRice    = "rice"
	CategoryNoodle  = "noodle"
	CategoryCoffee  = "coffee"
	CategorySnack   = "snack"
	CategoryMilkTea = "milk_tea"
	CategoryJuice   = "juice"
	CategoryChicken = "chicken"
	CategoryPizza   = "pizza"
	CategoryBurger  = "burger"
	CategoryPho     = "pho"
	CategoryBun     = "bun"
	CategoryBanhMi  = "banh_mi"
	CategoryOther   = "other"
)

func (*MenuItem) IsValidCategory(category string) bool {
	switch category {
	case CategoryRice, CategoryNoodle, CategoryCoffee, CategorySnack, CategoryMilkTea, CategoryJuice, CategoryChicken, CategoryPizza, CategoryBurger, CategoryPho, CategoryBun, CategoryBanhMi, CategoryOther:
		return true
	default:
		return false
	}
}

var (
	ErrInvalidCategory = errors.New("invalid category")
	ErrAlreadyFavorite = errors.New("already favorite")
)
