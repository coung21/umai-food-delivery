package repository

import (
	"context"
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
	if err := r.db.Model(olduser).Updates(upd).Scan(olduser).Error; err != nil {
		return nil, err
	}
	return olduser, nil
}
