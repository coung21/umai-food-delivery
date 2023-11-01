package usecase

import (
	"common"
	"context"
	"umai-auth-service/model"
)

func (u *authUC) RestaurantRegis(ctx context.Context, res *model.Restaurant) (*model.Restaurant, error) {
	user, err := u.authRepo.FindUserByID(ctx, res.UserID)
	if user == nil && err != nil {
		return nil, common.NotExistAccout
	}

	if err := u.authRepo.UpdateRole(ctx, user); err != nil {
		return nil, err
	}

	restaurant, err := u.authRepo.InsertRestaurant(ctx, res)

	if err != nil {
		return nil, err
	}
	return restaurant, nil
}
