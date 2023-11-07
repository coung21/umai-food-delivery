package usecase

import (
	"context"
	"umai-auth-service/model"
)

func (u *authUC) UpdateProfile(ctx context.Context, userid int, user *model.UserUpdate) (*model.User, error) {

}
