package model

import (
	"common"
	"time"
)

type Restaurant struct {
	common.SqlModel
	UserID         int        `json:"user_id" gorm:"column:user_id"`
	RestaurantName string     `json:"restaurant_name" gorm:"column:restaurant_name"`
	Slogan         string     `json:"slogan" gorm:"column:slogan"`
	OpenHour       *time.Time `json:"open_hour" gorm:"column:open_hour"`
	CloseHour      *time.Time `json:"close_hour" gorm:"column:close_hour"`
	Rating         float32    `json:"rating" gorm:"column:rating"`
}

func (Restaurant) TableName() string {
	return "restaurants"
}
