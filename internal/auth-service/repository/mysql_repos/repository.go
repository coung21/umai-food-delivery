package repository

import (
	"umai-auth-service/model"

	"gorm.io/gorm"
)

type authRepo struct {
	db *gorm.DB
}

func NewAuthRepo(db *gorm.DB) *authRepo {
	db.AutoMigrate(&model.User{})
	return &authRepo{db: db}
}
