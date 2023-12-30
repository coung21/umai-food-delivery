package repository

import (
	"gorm.io/gorm"
)

type menuRepoMysql struct {
	db *gorm.DB
}

func NewMenuRepo(db *gorm.DB) *menuRepoMysql {
	// db.AutoMigrate(&model.MenuItem{})
	return &menuRepoMysql{db: db}
}
