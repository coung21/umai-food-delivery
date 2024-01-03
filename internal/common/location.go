package common

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type Location struct {
	PlaceID  string  `json:"place_id" gorm:"column:place_id"` // PlaceID is the ID of the place in Goong
	Address  string  `json:"address" gorm:"column:address"`
	Lat      float64 `json:"lat" gorm:"column:lat"`
	Lng      float64 `json:"lng" gorm:"column:lng"`
	Compound struct {
		District string `json:"district"`
		Comune   string `json:"comune"`
		Province string `json:"province"`
	} `json:"compound"`
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
func (l *Location) Value() (driver.Value, error) {
	bytes, err := json.Marshal(l)
	if err != nil {
		return nil, err
	}

	return string(bytes), nil
}
