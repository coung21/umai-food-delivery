package menu

import (
	"context"
	"menu-service/model"
)

type Usecase interface {
	CreateMenuItem(ctx context.Context, mitem *model.MenuItem) (interface{}, error)
	ListMenuItemByResID(ctx context.Context, rid int) (*[]model.MenuItem, error)
	UpdateMenuItem(ctx context.Context, rid int, mid int, upd *model.UpdateMenuItem) (*model.MenuItem, error)
	DeleteMenuItem(ctx context.Context, mid int, rid int) (int, error)
	GetMenuItem(ctx context.Context, id int) (*model.MenuItem, error)
	AddFavorite(ctx context.Context, uid int, mid int) (*int, error)
	DeleteFavorite(ctx context.Context, uid int, mid int) (*int, error)
	ListFavorites(ctx context.Context, uid int) (*[]model.MenuItem, error)
}
