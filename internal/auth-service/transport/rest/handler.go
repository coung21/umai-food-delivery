package rest

import (
	auth "umai-auth-service/interfaces"
)

type authHandler struct {
	authUC auth.Usecase
}

func NewAuthHandler(authUC auth.Usecase) *authHandler {
	return &authHandler{authUC: authUC}
}
