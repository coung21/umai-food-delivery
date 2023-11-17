package menu

import (
	"context"
	"menu-service/model"
)

type Usecase interface {
	CreateMenuItem(ctx context.Context, mitem *model.MenuItem) (interface{}, error)
	ListMenuItemByResID(ctx context.Context, rid int) (*[]model.MenuItem, error)
}
