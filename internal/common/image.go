package common

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type Image struct {
	// Id     int    `json:"id" gorm:"column:id;"`
	Url       string `json:"url" gorm:"column:url;" bson:"url"`
	Width     int    `json:"width" gorm:"column:width;" bson:"width"`
	Height    int    `json:"height" gorm:"column:height;" bson:"height"`
	Ext       string `json:"ext" gorm:"column:ext" bson:"ext"`
	CloudName string `json:"cloud_name" gorm:"column:cloud_name" bson:"cloud_name"`
}

func (Image) TableName() string { return "images" }

func (j *Image) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	var img Image
	if err := json.Unmarshal(bytes, &img); err != nil {
		return err
	}

	*j = img
	return nil
}

// Value return json value, implement driver.Valuer interface
func (j *Image) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}
	return json.Marshal(j)
}

type Images []Image

func (j *Images) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	var img []Image
	if err := json.Unmarshal(bytes, &img); err != nil {
		return err
	}

	*j = img
	return nil
}

// Value return json value, implement driver.Valuer interface
func (j *Images) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}
	return json.Marshal(j)
}
