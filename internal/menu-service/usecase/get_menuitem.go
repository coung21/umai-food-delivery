package usecase

import (
	"context"
	"log"
	"menu-service/model"
	"time"
)

func (u *menuUC) GetMenuItem(ctx context.Context, id string) (*model.MenuItem, error) {
	// mitem, err := u.menuRepo.FindMenuItemByID(ctx, id)
	// if err != nil {
	// 	return nil, err
	// }
	// return mitem, nil
	cachedMitem, err := u.cacheRepo.Get(ctx, id)
	if err == nil {
		return cachedMitem, nil
	} else {
		log.Printf("Get Cache Err: %v", err)
	}

	mitem, err := u.menuRepo.FindMenuItemByID(ctx, id)
	if err != nil {
		return nil, err
	}

	err = u.cacheRepo.Set(ctx, id, mitem, time.Hour)
	if err != nil {
		log.Printf("Set Cache Err: %v", err)
	}

	return mitem, nil
}
