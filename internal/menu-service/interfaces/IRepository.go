package menu

import (
	"context"
	"menu-service/model"
)

type Repository interface {
	InsertMenuItem(ctx context.Context, mitem *model.MenuItem) (interface{}, error)
	ListMenuItemByResID(ctx context.Context, rid int) ([]model.MenuItem, error)
	UpdateMenuItem(ctx context.Context, rid int, mid string, upd *model.UpdateMenuItem) (*model.MenuItem, error)
	DeleteMenuItem(ctx context.Context, id string) (int, error)
	FindMenuItemByID(ctx context.Context, id string) (*model.MenuItem, error)
}
