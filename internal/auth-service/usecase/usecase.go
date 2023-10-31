package usecase

import (
	jwt "umai-auth-service/component"
	auth "umai-auth-service/interfaces"
)

type authUC struct {
	authRepo      auth.Repository
	tokenProvider jwt.TokenProvider
	expToken      int
}

func NewAuthUC(repo auth.Repository, tokenprovider jwt.TokenProvider, exp int) *authUC {
	return &authUC{authRepo: repo, tokenProvider: tokenprovider, expToken: exp}
}
