package menu

import (
	"context"
	"menu-service/model"
)

type Usecase interface {
	CreateMenuItem(ctx context.Context, mitem *model.MenuItem) (interface{}, error)
	ListMenuItemByResID(ctx context.Context, rid int) (*[]model.MenuItem, error)
	UpdateMenuItem(ctx context.Context, rid int, mid string, upd *model.UpdateMenuItem) (*model.MenuItem, error)
	DeleteMenuItem(ctx context.Context, mid string, rid int) (int, error)
	GetMenuItem(ctx context.Context, id string) (*model.MenuItem, error)
}
