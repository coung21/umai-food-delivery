package repository

import "gorm.io/gorm"

type menuRepo struct {
	db *gorm.DB
}

func NewMenuRepo(db *gorm.DB) *menuRepo {
	return &menuRepo{db: db}
}
