package usecase

import (
	"common"
	"context"
	"umai-auth-service/model"
)

func (u *authUC) UpdateProfile(ctx context.Context, userid int, udp *model.UserUpdate) (*model.User, error) {
	oldData, err := u.authRepo.FindUserByID(ctx, userid)
	if oldData == nil && err != nil {
		return nil, common.NotExistAccount
	}

	newData, err := u.authRepo.UpdateUser(ctx, oldData, udp)
	if err != nil {
		return nil, common.InternalServerError
	}

	newData.SanitizePassword()
	return newData, nil
}
