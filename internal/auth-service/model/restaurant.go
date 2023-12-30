package model

import (
	"common"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type Restaurant struct {
	common.SqlModel
	UserID         int            `json:"user_id" gorm:"column:user_id;not null;index"`
	RestaurantName string         `json:"restaurant_name" gorm:"column:restaurant_name;not null"`
	Slogan         string         `json:"slogan" gorm:"column:slogan"`
	Cover          *common.Images `json:"cover" gorm:"column:cover;type:json"`
	OpenHour       *time.Time     `json:"open_hour" gorm:"column:open_hour"`
	CloseHour      *time.Time     `json:"close_hour" gorm:"column:close_hour"`
	Location       Location       `json:"location" gorm:"column:location;null;type:json"`
	Rating         float32        `json:"rating" gorm:"column:rating"`
}

func (Restaurant) TableName() string {
	return "restaurants"
}

type RestaurantUpdate struct {
	RestaurantName *string        `json:"restaurant_name" gorm:"column:restaurant_name;type:varchar(255);"`
	Slogan         *string        `json:"slogan" gorm:"column:slogan;type:text;"`
	Cover          *common.Images `json:"cover" gorm:"column:cover;type:json"`
	OpenHour       *time.Time     `json:"open_hour" gorm:"column:open_hour;type:time;"`
	CloseHour      *time.Time     `json:"close_hour" gorm:"column:close_hour;type:time;"`
	Location       *Location      `json:"location" gorm:"column:location"`
}

type Location struct {
	Address string  `json:"address" gorm:"column:address"`
	Lat     float64 `json:"lat" gorm:"column:lat"`
	Lng     float64 `json:"lng" gorm:"column:lng"`
}

// Implement the scanner interface for Location
func (l *Location) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("failed to scan Location: invalid data type")
	}

	return json.Unmarshal(bytes, l)
}

// Implement the valuer interface for Location
func (l Location) Value() (driver.Value, error) {
	bytes, err := json.Marshal(l)
	if err != nil {
		return nil, err
	}

	return string(bytes), nil
}
