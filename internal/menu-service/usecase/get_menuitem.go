package usecase

import (
	"context"
	"menu-service/model"
)

func (u *menuUC) GetMenuItem(ctx context.Context, id string) (*model.MenuItem, error) {
	mitem, err := u.menuRepo.FindMenuItemByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return mitem, nil
}
