package repository

import (
	"context"
	"umai-auth-service/model"
)

func (r *authRepo) FindUserByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *authRepo) FindUserByID(ctx context.Context, id int) (*model.User, error) {
	var user model.User

	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *authRepo) FindRestaurantByID(ctx context.Context, id int) (*model.Restaurant, error) {
	var res model.Restaurant

	if err := r.db.Where("id = ?", id).First(&res).Error; err != nil {
		return nil, err
	}
	return &res, nil
}

func (r *authRepo) FindRestaurantByUserID(ctx context.Context, uid int) (*model.Restaurant, error) {
	var res model.Restaurant
	if err := r.db.Where("user_id = ?", uid).First(&res).Error; err != nil {
		return nil, err
	}
	return &res, nil
}
