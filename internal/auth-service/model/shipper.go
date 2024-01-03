package model

import "time"

type Shipper struct {
	ID           int    `json:"id" gorm:"column:id;primary_key;auto_increment"`
	UserID       int    `json:"user_id" gorm:"column:user_id;not null;index"`
	VehicleName  string `json:"vehicle_name" gorm:"column:vehicle_name;not null"`
	VehicleType  string `json:"vehicle_type" gorm:"column:vehicle_type;not null"`
	LicensePlate string `json:"license_plate" gorm:"column:license_plate;not null"`
	WorkZone     struct {
		Province string `json:"province" gorm:"column:province;not null"`
		District string `json:"district" gorm:"column:district;not null"`
	} `json:"work_zone" gorm:"column:work_zone;not null;type:json"`
	WorkStatus string    `json:"work_status" gorm:"column:work_status;not null"`
	UpdateAt   time.Time `json:"update_at" gorm:"column:update_at;not null"`
	CreateAt   time.Time `json:"create_at" gorm:"column:create_at;not null"`
}

func (Shipper) TableName() string {
	return "shippers"
}

type ShipperUpdate struct {
	VehicleName  *string `json:"vehicle_name" gorm:"column:vehicle_name;not null"`
	VehicleType  *string `json:"vehicle_type" gorm:"column:vehicle_type;not null"`
	LicensePlate *string `json:"license_plate" gorm:"column:license_plate;not null"`
	WorkZone     struct {
		Province *string `json:"province" gorm:"column:province;not null"`
		District *string `json:"district" gorm:"column:district;not null"`
	} `json:"work_zone" gorm:"column:work_zone;not null;type:json"`
	WorkStatus *string `json:"work_status" gorm:"column:work_status;not null"`
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

func (s *Shipper) IsInWorkZone(province, district string) bool {
	return s.WorkZone.Province == province && s.WorkZone.District == district
}
