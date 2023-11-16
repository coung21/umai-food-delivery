package usecase

import (
	"common"
	"context"
	"umai-auth-service/model"
)

func (u *authUC) Register(ctx context.Context, user *model.User) (int, error) {
	if existUser, err := u.authRepo.FindUserByEmail(ctx, user.Email); existUser != nil && err == nil {
		return 0, common.ExistsEmailError
	}

	if err := user.HashPassword(); err != nil {
		return 0, common.InternalServerError
	}

	user.DefaultRole()

	createdUserID, err := u.authRepo.InsertUser(ctx, user)
	if err != nil {
		return 0, err
	}

	return createdUserID, nil

}
