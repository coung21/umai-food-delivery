package auth

import (
	"context"
	"umai-auth-service/model"
)

type Repository interface {
	FindUserByEmail(ctx context.Context, email string) (*model.User, error)
	InsertUser(ctx context.Context, user *model.User) (*model.User, error)
}
