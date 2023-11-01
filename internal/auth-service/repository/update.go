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
