package repository

import (
	"context"
	"menu-service/model"
)

func (r *menuRepoMysql) DeleteMenuItem(ctx context.Context, id int) (int, error) {
	var mitem model.MenuItem
	err := r.db.Where("id = ?", id).Delete(&mitem).Error
	if err != nil {
		return 0, err
	}
	return id, nil
}
