package rest

import (
	jwt "umai-auth-service/component"
	auth "umai-auth-service/interfaces"
)

type authHandler struct {
	authUC        auth.Usecase
	tokenProvider jwt.TokenProvider
	authRepo      auth.Repository
}

func NewAuthHandler(authUC auth.Usecase, authRepo auth.Repository, tokenprovider jwt.TokenProvider) *authHandler {
	return &authHandler{authUC: authUC, authRepo: authRepo, tokenProvider: tokenprovider}
}
