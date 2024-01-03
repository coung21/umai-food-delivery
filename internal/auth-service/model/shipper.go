package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type Shipper struct {
	ID           int       `json:"id" gorm:"column:id;primary_key;auto_increment"`
	UserID       int       `json:"user_id" gorm:"column:user_id;not null;index"`
	VehicleName  string    `json:"vehicle_name" gorm:"column:vehicle_name;not null"`
	VehicleType  string    `json:"vehicle_type" gorm:"column:vehicle_type;not null"`
	LicensePlate string    `json:"license_plate" gorm:"column:license_plate;not null"`
	WorkStatus   string    `json:"work_status" gorm:"column:work_status;not null"`
	WorkZone     WorkZone  `json:"work_zone" gorm:"column:work_zone;not null;type:json"`
	UpdateAt     time.Time `json:"update_at" gorm:"column:update_at;not null"`
	CreateAt     time.Time `json:"create_at" gorm:"column:create_at;not null"`
}

func (Shipper) TableName() string {
	return "shippers"
}

type ShipperUpdate struct {
	VehicleName  *string   `json:"vehicle_name" gorm:"column:vehicle_name;not null"`
	VehicleType  *string   `json:"vehicle_type" gorm:"column:vehicle_type;not null"`
	LicensePlate *string   `json:"license_plate" gorm:"column:license_plate;not null"`
	WorkStatus   *string   `json:"work_status" gorm:"column:work_status;not null"`
	WorkZone     *WorkZone `json:"work_zone" gorm:"column:work_zone;not null;type:json"`
}

type WorkZone struct {
	Province string `json:"province" gorm:"column:province;not null"`
	District string `json:"district" gorm:"column:district;not null"`
}

func (s *Shipper) UpdateStatus(status string) {
	s.WorkStatus = status
}

const (
	ShipperStatusAvailable = "available"
	ShipperStatusBusy      = "busy"
)

const (
	ShipperVehicleTypeMotorbike = "motorbike"
	ShipperVehicleTypeBicycle   = "bicycle"
)

func (s *Shipper) IsAvailable() bool {
	return s.WorkStatus == ShipperStatusAvailable
}

func (s *Shipper) IsBusy() bool {
	return s.WorkStatus == ShipperStatusBusy
}

func (s *Shipper) IsMotorbike() bool {
	return s.VehicleType == ShipperVehicleTypeMotorbike
}

func (s *Shipper) IsBicycle() bool {
	return s.VehicleType == ShipperVehicleTypeBicycle
}

func (s *WorkZone) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	var workZone WorkZone
	if err := json.Unmarshal(bytes, &workZone); err != nil {
		return err
	}
	*s = workZone
	return nil
}

func (s *WorkZone) Value() (driver.Value, error) {
	if s == nil {
		return nil, nil
	}
	return json.Marshal(s)
}
