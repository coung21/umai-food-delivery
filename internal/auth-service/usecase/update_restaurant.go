package usecase

import (
	"common"
	"context"
	"umai-auth-service/model"
)

func (u *authUC) UpdateRestaurant(ctx context.Context, id int, udp *model.RestaurantUpdate) (*model.Restaurant, error) {
	cuserId := ctx.Value(common.CuserId)
	olddata, err := u.authRepo.FindRestaurantByID(ctx, id)
	if err != nil && olddata == nil {
		return nil, common.NotExistAccount
	}

	if olddata.UserID != cuserId {
		return nil, common.Forbidden
	}

	newData, err := u.authRepo.UpdateRestaurant(ctx, olddata, udp)
	if err != nil {
		return nil, common.InternalServerError
	}

	return newData, nil
}
