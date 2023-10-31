package usecase

import (
	"common"
	"context"
	"umai-auth-service/model"
)

func (u *authUC) Login(ctx context.Context, cred *model.LoginCredentials) (*model.UserWithToken, error) {
	user, err := u.authRepo.FindUserByEmail(ctx, cred.Email)

	if err != nil && user == nil {
		return nil, common.NotExistAccout
	}

	isMatch := user.ComparePassword(cred.Password)

	if !isMatch {
		return nil, common.WrongPassword
	}

	payload := u.tokenProvider.NewPayLoad(int(user.ID), string(user.Role))

	token, err := u.tokenProvider.GenerateToken(payload, u.expToken)
	if err != nil {
		return nil, common.InvalidJWTClaims
	}

	return &model.UserWithToken{
		User:  *user,
		Token: *token,
	}, nil
}
