package usecase

import (
	"umai-auth-service/common"
	jwt "umai-auth-service/component"
	auth "umai-auth-service/interfaces"
	"umai-auth-service/model"
)

type authUC struct {
	authRepo      auth.Repository
	tokenProvider jwt.TokenProvider
}

func NewAuthUC(repo auth.Repository, tokenprovider jwt.TokenProvider) *authUC {
	return &authUC{authRepo: repo, tokenProvider: tokenprovider}
}

func (u *authUC) Register(user *model.User) (*model.User, error) {
	if existUser, err := u.authRepo.FindUserByEmail(user.Email); existUser != nil && err == nil {
		return nil, common.ExistsEmailError
	}

	if err := user.HashPassword(); err != nil {
		return nil, common.InternalServerError
	}

	createdUser, err := u.authRepo.InsertUser(user)
	if err != nil {
		return nil, err
	}

	createdUser.EncodeId()
	createdUser.SanitizePassword()

	return createdUser, nil

}
