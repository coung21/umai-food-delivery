package usecase

import (
	"common"
	"context"
	"umai-auth-service/model"
)

func (u *authUC) UpdateRestaurant(ctx context.Context, id int, udp *model.RestaurantUpdate) (*model.Restaurant, error) {
	cuser := ctx.Value(common.CurrentUser).(*model.User)
	olddata, err := u.authRepo.FindRestaurantByID(ctx, id)
	if err != nil && olddata == nil {
		return nil, common.NotExistAccount
	}

	if olddata.UserID != cuser.ID {
		return nil, common.Forbidden
	}

	newData, err := u.authRepo.UpdateRestaurant(ctx, olddata, udp)
	if err != nil {
		return nil, common.InternalServerError
	}

	return newData, nil
}
