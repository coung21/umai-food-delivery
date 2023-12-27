package usecase

import (
	"context"
	"menu-service/model"
)

func (u *menuUC) ListMenuItemsByCategory(ctx context.Context, category string) (*[]model.MenuItem, error) {
	menuItems, err := u.menuRepo.ListMenuItemByCategory(ctx, category)
	if err != nil {
		return nil, err
	}

	return &menuItems, nil
}
