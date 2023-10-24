package auth

import "umai-auth-service/model"

type Repository interface {
	FindUserByEmail(email string) (*model.User, error)
	InsertUser(user *model.User) (*model.User, error)
}
