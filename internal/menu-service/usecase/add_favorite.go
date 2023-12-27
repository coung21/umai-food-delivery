package usecase

import (
	"context"
	"menu-service/model"
)

func (u *menuUC) AddFavorite(ctx context.Context, uid int, mid string) (*string, error) {
	// 1. check if menu is exist
	mitem, err := u.menuRepo.FindMenuItemByID(ctx, mid)
	if err != nil {
		return nil, err
	}
	// 2. check if menu is already favorite
	if ex, _ := u.cacheRepo.GetFavorite(ctx, uid, mid); !ex {
		return nil, model.ErrAlreadyFavorite
	}

	// 3. add favorite
	if err = u.cacheRepo.SetFavorite(ctx, uid, mid); err != nil {
		return nil, err
	}
	// 4. return favorite id
	favoriteID := mitem.ID.Hex() // Convert ObjectID to string
	return &favoriteID, nil
}
