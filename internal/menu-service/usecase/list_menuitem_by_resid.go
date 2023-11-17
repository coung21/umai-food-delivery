package usecase

import (
	"context"
	"menu-service/model"
)

func (u *menuUC) ListMenuItemByResID(ctx context.Context, rid int) (*[]model.MenuItem, error) {
	data, err := u.menuRepo.ListMenuItemByResID(ctx, rid)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
