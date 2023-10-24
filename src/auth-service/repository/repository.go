package repository

import (
	"umai-auth-service/model"

	"gorm.io/gorm"
)

type authRepo struct {
	db *gorm.DB
}

func NewAuthRepo(db *gorm.DB) *authRepo {
	return &authRepo{db: db}
}

func (r *authRepo) FindUserByEmail(email string) (*model.User, error) {
	var user model.User
	result := r.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *authRepo) InsertUser(user *model.User) (*model.User, error) {
	result := r.db.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}
