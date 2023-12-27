package usecase

import (
	"context"
	"menu-service/model"
)

func (u *menuUC) CreateMenuItem(ctx context.Context, mitem *model.MenuItem) (interface{}, error) {
	if !mitem.IsValidCategory(mitem.Category) {
		return nil, model.ErrInvalidCategory
	}

	newMItemId, err := u.menuRepo.InsertMenuItem(ctx, mitem)
	if err != nil {
		return nil, err
	}
	return newMItemId, nil
}
