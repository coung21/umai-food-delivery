package menu

import (
	"context"
	"menu-service/model"
)

type Usecase interface {
	CreateMenuItem(ctx context.Context, mitem *model.MenuItem) (interface{}, error)
}
