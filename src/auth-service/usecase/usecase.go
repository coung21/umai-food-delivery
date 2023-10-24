package usecase

import auth "umai-auth-service/interfaces"

type authUC struct {
	authRepo auth.Repository
}

func NewAuthUC(repo auth.Repository) *authUC {
	return &authUC{authRepo: repo}
}
