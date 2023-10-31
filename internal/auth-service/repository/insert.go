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
