package usecase

import (
	"context"
	"menu-service/model"
)

func (u *menuUC) ListFavorites(ctx context.Context, uid int) (*[]model.MenuItem, error) {
	// 1. get favorites id from cache
	favorites, err := u.cacheRepo.ListFavorites(ctx, uid)
	if err != nil {
		return nil, err
	}

	// 2. get menu items from find menu items by id in for loop
	var menuItems []model.MenuItem
	for _, favorite := range favorites {
		menuItem, err := u.menuRepo.FindMenuItemByID(ctx, favorite)
		if err != nil {
			return nil, err
		}
		menuItems = append(menuItems, *menuItem)
	}
	return &menuItems, nil
}
