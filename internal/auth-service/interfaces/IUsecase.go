package auth

import (
	"context"
	"umai-auth-service/model"
)

type Usecase interface {
	Register(ctx context.Context, user *model.User) (*model.User, error)
	Login(ctx context.Context, cred *model.LoginCredentials) (*model.UserWithToken, error)
	RestaurantRegis(ctx context.Context, res *model.Restaurant) (*model.Restaurant, error)
	UpdateProfile(ctx context.Context, userid int, udp *model.UserUpdate) (*model.User, error)
	GetProfile(ctx context.Context, id int) (*model.User, error)
	// GetRestaurent(ctx context.Context, id int) (*model.Restaurant, error)
}
