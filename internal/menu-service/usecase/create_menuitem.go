package usecase

import (
	"context"
	"menu-service/model"
)

func (u *menuUC) CreateMenuItem(ctx context.Context, mitem *model.MenuItem) (interface{}, error) {
	newMItemId, err := u.menuRepo.InsertMenuItem(ctx, mitem)
	if err != nil {
		return 0, err
	}
	return newMItemId, nil
}
