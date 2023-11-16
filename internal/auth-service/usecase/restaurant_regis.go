package usecase

import (
	"common"
	"context"
	"umai-auth-service/model"
)

func (u *authUC) RestaurantRegis(ctx context.Context, res *model.Restaurant) (int, error) {
	user, err := u.authRepo.FindUserByID(ctx, res.UserID)
	if user == nil && err != nil {
		return 0, common.NotExistAccount
	}

	if user.Role == model.RoleCustomer {
		if err := u.authRepo.UpdateRole(ctx, user); err != nil {
			return 0, err
		}
	} else {
		return 0, common.BadRequest
	}

	restaurantId, err := u.authRepo.InsertRestaurant(ctx, res)

	if err != nil {
		return 0, err
	}
	return restaurantId, nil
}
