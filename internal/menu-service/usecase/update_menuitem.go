package usecase

import (
	"context"
	"menu-service/model"
)

func (u *menuUC) UpdateMenuItem(ctx context.Context, rid int, mid string, upd *model.UpdateMenuItem) (*model.MenuItem, error) {
	mitem, err := u.menuRepo.UpdateMenuItem(ctx, rid, mid, upd)
	if err != nil {
		return nil, err
	}
	return mitem, nil
}
