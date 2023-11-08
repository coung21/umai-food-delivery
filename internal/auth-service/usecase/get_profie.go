package usecase

import (
	"common"
	"context"
	"umai-auth-service/model"
)

func (u *authUC) GetProfile(ctx context.Context, id int) (*model.User, error) {
	user, err := u.authRepo.FindUserByID(ctx, id)
	if err != nil && user == nil {
		return nil, common.NotExistAccount
	}

	user.SanitizePassword()
	return user, nil
}
