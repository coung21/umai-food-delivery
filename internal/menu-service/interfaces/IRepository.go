package menu

import (
	"context"
	"menu-service/model"
)

type Repository interface {
	InsertMenuItem(ctx context.Context, mitem *model.MenuItem) (interface{}, error)
	ListMenuItemByResID(ctx context.Context, rid int) ([]model.MenuItem, error)
}
