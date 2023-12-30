package repository

import "gorm.io/gorm"

type menuRepoMysql struct {
	db *gorm.DB
}

func NewMenuRepo(db *gorm.DB) *menuRepoMysql {
	return &menuRepoMysql{db: db}
}
