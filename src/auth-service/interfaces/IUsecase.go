package auth

import "umai-auth-service/model"

type Usecase interface {
	Register(user *model.User) (*model.User, error)
}
