package model

import (
	"common"
	"time"
)

type Restaurant struct {
	common.SqlModel
	UserID         int              `json:"user_id" gorm:"column:user_id;not null;index"`
	RestaurantName string           `json:"restaurant_name" gorm:"column:restaurant_name;not null"`
	Slogan         string           `json:"slogan" gorm:"column:slogan"`
	Cover          *common.Images   `json:"cover" gorm:"column:cover;type:json"`
	OpenHour       *time.Time       `json:"open_hour" gorm:"column:open_hour"`
	CloseHour      *time.Time       `json:"close_hour" gorm:"column:close_hour"`
	Location       *common.Location `json:"location" gorm:"column:location;null;type:json"`
	Rating         float32          `json:"rating" gorm:"column:rating"`
}

func (Restaurant) TableName() string {
	return "restaurants"
}

type RestaurantUpdate struct {
	RestaurantName *string          `json:"restaurant_name" gorm:"column:restaurant_name;type:varchar(255);"`
	Slogan         *string          `json:"slogan" gorm:"column:slogan;type:text;"`
	Cover          *common.Images   `json:"cover" gorm:"column:cover;type:json"`
	OpenHour       *time.Time       `json:"open_hour" gorm:"column:open_hour;type:time;"`
	CloseHour      *time.Time       `json:"close_hour" gorm:"column:close_hour;type:time;"`
	Location       *common.Location `json:"location" gorm:"column:location"`
}
