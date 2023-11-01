package repository

import (
	"context"
	"umai-auth-service/model"
)

func (r *authRepo) InsertUser(ctx context.Context, user *model.User) (*model.User, error) {
	if err := r.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *authRepo) InsertRestaurant(ctx context.Context, res *model.Restaurant) (*model.Restaurant, error) {
	db := r.db.Begin()

	if err := db.Table(res.TableName()).Create(res).Error; err != nil {
		db.Rollback()
		return nil, err
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return nil, err
	}

	return res, nil
}
