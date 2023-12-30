package model

import (
	"common"
	"time"
)

type Discount struct {
	common.SqlModel
	RestaurantID int        `json:"restaurant_id" gorm:"column:restaurant_id"`
	DiscountName string     `json:"discount_name" gorm:"column:discount_name"`
	DiscountType string     `json:"discount_type" gorm:"column:discount_type"`
	Discount     float32    `json:"discount" gorm:"column:discount"`
	StartDate    *time.Time `json:"start_date" gorm:"column:start_date"`
	EndDate      *time.Time `json:"end_date" gorm:"column:end_date"`
}
