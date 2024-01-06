package repository

import (
	"context"
	"time"
	"umai-auth-service/model"
)

func (r *authRepo) UpdateRole(ctx context.Context, user *model.User) error {
	db := r.db.Begin()

	if err := db.Table(user.TableName()).Model(user).Update("role", model.RoleRestaurant).Error; err != nil {
		db.Rollback()
		return err
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return err
	}
	return nil
}

func (r *authRepo) UpdateUser(ctx context.Context, olduser *model.User, upd *model.UserUpdate) (*model.User, error) {
	upd.UpdateAt = time.Now()
	if err := r.db.Model(olduser).Updates(upd).Scan(olduser).Error; err != nil {
		return nil, err
	}
	return olduser, nil
}

func (r *authRepo) UpdateRestaurant(ctx context.Context, oldres *model.Restaurant, upd *model.RestaurantUpdate) (*model.Restaurant, error) {
	if err := r.db.Model(oldres).Updates(upd).Scan(oldres).Error; err != nil {
		return nil, err
	}
	return oldres, nil
}

func (r *authRepo) UpdateShipper(ctx context.Context, oldshipper *model.Shipper, upd *model.ShipperUpdate) (*model.Shipper, error) {
	if err := r.db.Model(oldshipper).Updates(upd).Scan(oldshipper).Error; err != nil {
		return nil, err
	}
	return oldshipper, nil
}
