package usecase

import (
	"common"
	"context"
)

func (u *menuUC) DeleteMenuItem(ctx context.Context, mid string, rid int) (int, error) {
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
	return delCount, nil
}
