package usecase

import (
	"common"
	"context"
	"log"
)

func (u *menuUC) DeleteMenuItem(ctx context.Context, mid int, rid int) (int, error) {
	crid := ctx.Value(common.CResId).(int)
	mitem, err := u.menuRepo.FindMenuItemByID(ctx, mid)
	if err != nil {
		return 0, err
	}
	if rid != mitem.RestaurantID || crid != mitem.RestaurantID {
		return 0, common.Forbidden
	}
	delCount, err := u.menuRepo.DeleteMenuItem(ctx, mid)
	if err != nil {
		return delCount, err
	}

	if err := u.cacheRepo.Del(ctx, mid); err != nil {
		log.Printf("Del Cache Err : %v", err)
	}
	return delCount, nil
}
