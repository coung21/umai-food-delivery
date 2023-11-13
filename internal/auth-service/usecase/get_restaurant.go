package usecase

import (
	"common"
	"context"
	"umai-auth-service/model"
)

func (u *authUC) GetRestaurant(ctx context.Context, id int) (*model.Restaurant, error) {
	res, err := u.authRepo.FindRestaurantByID(ctx, id)
	if err != nil && res == nil {
		return nil, common.NotExistAccount
	}

	return res, nil
}
