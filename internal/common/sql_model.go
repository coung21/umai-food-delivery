package common

import "time"

type SqlModel struct {
	ID        int       `json:"id" gorm:"column:id"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at,omitempty"`
	CreatedAt time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at,omitempty"`
}
