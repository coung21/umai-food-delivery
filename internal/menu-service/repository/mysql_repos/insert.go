package repository

import (
	"context"
	"menu-service/model"
)

func (r *menuRepoMysql) InsertMenuItem(ctx context.Context, mitem *model.MenuItem) (interface{}, error) {
	result := r.db.Create(&mitem)
	if result.Error != nil {
		return nil, result.Error
	}
	return mitem.ID, nil
}
