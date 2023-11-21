package usecase

import (
	jwt "umai-auth-service/component"
	auth "umai-auth-service/interfaces"
)

type authUC struct {
	authRepo      auth.Repository
	cacheRepo     auth.CacheRepository
	tokenProvider jwt.TokenProvider
	expToken      int
}

func NewAuthUC(repo auth.Repository, cacheRepo auth.CacheRepository, tokenprovider jwt.TokenProvider, exp int) *authUC {
	return &authUC{authRepo: repo, cacheRepo: cacheRepo, tokenProvider: tokenprovider, expToken: exp}
}
