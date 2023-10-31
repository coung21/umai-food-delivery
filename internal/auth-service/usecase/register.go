package usecase

import (
	"common"
	"context"
	"umai-auth-service/model"
)

func (u *authUC) Register(ctx context.Context, user *model.User) (*model.User, error) {
	if existUser, err := u.authRepo.FindUserByEmail(ctx, user.Email); existUser != nil && err == nil {
		return nil, common.ExistsEmailError
	}

	if err := user.HashPassword(); err != nil {
		return nil, common.InternalServerError
	}

	user.Role.Default()

	createdUser, err := u.authRepo.InsertUser(ctx, user)
	if err != nil {
		return nil, err
	}

	createdUser.EncodeId()
	createdUser.SanitizePassword()

	return createdUser, nil

}
