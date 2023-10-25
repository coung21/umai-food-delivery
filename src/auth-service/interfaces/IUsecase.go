package auth

import (
	"context"
	"umai-auth-service/model"
)

type Usecase interface {
	Register(ctx context.Context, user *model.User) (*model.User, error)
}
