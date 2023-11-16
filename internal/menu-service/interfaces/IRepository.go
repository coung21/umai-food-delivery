package menu

import (
	"context"
	"menu-service/model"
)

type Repository interface {
	InsertMenuItem(ctx context.Context, mitem *model.MenuItem) (*model.MenuItem, error)
}
