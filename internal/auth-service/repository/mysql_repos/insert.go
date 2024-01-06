package repository

import (
	"context"
	"umai-auth-service/model"
)

func (r *authRepo) InsertUser(ctx context.Context, user *model.User) (int, error) {
	result := r.db.Create(user)
	if result.Error != nil {
		return 0, result.Error
	}
	return user.ID, nil
}

func (r *authRepo) InsertRestaurant(ctx context.Context, res *model.Restaurant) (int, error) {
	db := r.db.Begin()

	result := db.Table(res.TableName()).Create(res)

	if result.Error != nil {
		db.Rollback()
		return 0, result.Error
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return 0, err
	}

	return res.ID, nil
}

func (r *authRepo) InsertShipper(ctx context.Context, shipper *model.Shipper) (int, error) {
	db := r.db.Begin()

	result := db.Table(shipper.TableName()).Create(shipper)

	if result.Error != nil {
		db.Rollback()
		return 0, result.Error
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return 0, err
	}

	return shipper.ID, nil
}
