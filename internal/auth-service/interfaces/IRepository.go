package auth

import (
	"context"
	"time"
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
	FindRestaurantByUserID(ctx context.Context, uid int) (*model.Restaurant, error)
	InsertShipper(ctx context.Context, shipper *model.Shipper) (int, error)
	UpdateShipper(ctx context.Context, oldshipper *model.Shipper, upd *model.ShipperUpdate) (*model.Shipper, error)
	FindShipperByID(ctx context.Context, id int) (*model.Shipper, error)
	FindShipperByUserID(ctx context.Context, uid int) (*model.Shipper, error)
}

type CacheRepository interface {
	Set(ctx context.Context, id int, res *model.Restaurant, ttl time.Duration) error
	Get(ctx context.Context, id int) (*model.Restaurant, error)
	Del(ctx context.Context, id int) error
}
