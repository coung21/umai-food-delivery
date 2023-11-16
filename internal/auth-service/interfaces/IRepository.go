package auth

import (
	"context"
	"umai-auth-service/model"
)

type Repository interface {
	FindUserByEmail(ctx context.Context, email string) (*model.User, error)
	InsertUser(ctx context.Context, user *model.User) (int, error)
	FindUserByID(ctx context.Context, id int) (*model.User, error)
	UpdateRole(ctx context.Context, user *model.User) error
	InsertRestaurant(ctx context.Context, res *model.Restaurant) (int, error)
	UpdateUser(ctx context.Context, olduser *model.User, upd *model.UserUpdate) (*model.User, error)
	FindRestaurantByID(ctx context.Context, id int) (*model.Restaurant, error)
	UpdateRestaurant(ctx context.Context, oldres *model.Restaurant, upd *model.RestaurantUpdate) (*model.Restaurant, error)
}
