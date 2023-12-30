package menu

import (
	"context"
	"menu-service/model"
	"time"
)

type Repository interface {
	InsertMenuItem(ctx context.Context, mitem *model.MenuItem) (interface{}, error)
	ListMenuItemByResID(ctx context.Context, rid int) ([]model.MenuItem, error)
	UpdateMenuItem(ctx context.Context, rid int, mid int, upd *model.UpdateMenuItem) (*model.MenuItem, error)
	DeleteMenuItem(ctx context.Context, id int) (int, error)
	FindMenuItemByID(ctx context.Context, id int) (*model.MenuItem, error)
	ListMenuItemByCategory(ctx context.Context, category string) ([]model.MenuItem, error)
}

type CacheRepository interface {
	Set(ctx context.Context, id int, mitem *model.MenuItem, ttl time.Duration) error
	Get(ctx context.Context, id int) (*model.MenuItem, error)
	Del(ctx context.Context, id int) error
	SetFavorite(ctx context.Context, uid int, mid int) error
	ListFavorites(ctx context.Context, uid int) ([]int, error)
	DelFavorite(ctx context.Context, uid int, mid int) error
	GetFavorite(ctx context.Context, uid int, mid int) (bool, error)
}
