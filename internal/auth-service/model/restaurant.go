package model

import (
	"common"
	"time"
)

type Restaurant struct {
	common.SqlModel
	UserID         int            `json:"user_id" gorm:"column:user_id"`
	RestaurantName string         `json:"restaurant_name" gorm:"column:restaurant_name"`
	Slogan         string         `json:"slogan" gorm:"column:slogan"`
	Cover          *common.Images `json:"cover" gorm:"column:cover"`
	OpenHour       *time.Time     `json:"open_hour" gorm:"column:open_hour"`
	CloseHour      *time.Time     `json:"close_hour" gorm:"column:close_hour"`
	Location       Location       `json:"location" gorm:"column:location"`
	Rating         float32        `json:"rating" gorm:"column:rating"`
}

func (Restaurant) TableName() string {
	return "restaurants"
}

type RestaurantUpdate struct {
	RestaurantName *string        `json:"restaurant_name" gorm:"column:restaurant_name"`
	Slogan         *string        `json:"slogan" gorm:"column:slogan"`
	Cover          *common.Images `json:"cover" gorm:"column:cover"`
	OpenHour       *time.Time     `json:"open_hour" gorm:"column:open_hour"`
	CloseHour      *time.Time     `json:"close_hour" gorm:"column:close_hour"`
	Location       *Location      `json:"location" gorm:"column:location"`
}

type Location struct {
	Address string  `json:"address" gorm:"column:address"`
	Lat     float64 `json:"lat" gorm:"column:lat"`
	Lng     float64 `json:"lng" gorm:"column:lng"`
}
