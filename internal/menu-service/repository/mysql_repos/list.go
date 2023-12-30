package repository

import (
	"context"
	"menu-service/model"
)

func (r *menuRepoMysql) ListMenuItemByResID(ctx context.Context, rid int) ([]model.MenuItem, error) {
	var menuitems []model.MenuItem
	err := r.db.Where("restaurant_id = ?", rid).Find(&menuitems).Error
	if err != nil {
		return nil, err
	}
	return menuitems, nil
}

func (r *menuRepoMysql) ListMenuItemByCategory(ctx context.Context, category string) ([]model.MenuItem, error) {
	var menuitems []model.MenuItem
	err := r.db.Where("category = ?", category).Find(&menuitems).Error
	if err != nil {
		return nil, err
	}
	return menuitems, nil
}
