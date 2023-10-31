package usecase

import (
	"common"
	"context"
	"log"
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

	payload := u.tokenProvider.NewPayLoad(int(user.ID), user.Role)

	token, err := u.tokenProvider.GenerateToken(payload, u.expToken)
	if err != nil {
		log.Fatal(err)
		return nil, common.InvalidJWTClaims
	}

	user.EncodeId()
	user.SanitizePassword()

	return &model.UserWithToken{
		User:  *user,
		Token: *token,
	}, nil
}
