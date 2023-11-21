package usecase

import (
	"common"
	"context"
	"log"
	"time"
	"umai-auth-service/model"
)

func (u *authUC) GetRestaurant(ctx context.Context, id int) (*model.Restaurant, error) {
	// res, err := u.authRepo.FindRestaurantByID(ctx, id)
	// if err != nil && res == nil {
	// 	return nil, common.NotExistAccount
	// }

	// return res, nil

	cachedRes, err := u.cacheRepo.Get(ctx, id)
	if err == nil {
		return cachedRes, nil
	} else {
		log.Printf("Get Cache Err: %v", err)
	}

	res, err := u.authRepo.FindRestaurantByID(ctx, id)
	if err != nil && res == nil {
		return nil, common.NotExistAccount
	}

	err = u.cacheRepo.Set(ctx, id, res, time.Hour)
	if err != nil {
		log.Printf("Set Cache Err: %v", err)
	}

	return res, nil
}
