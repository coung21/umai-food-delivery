package repository

import (
	"context"
	"menu-service/model"
	"time"
)

func (r *menuRepoMysql) UpdateMenuItem(ctx context.Context, rid int, mid int, upd *model.UpdateMenuItem) (*model.MenuItem, error) {
	var mitem model.MenuItem
	upd.UpdatedAt = time.Now()
	err := r.db.Model(&mitem).Where("restaurant_id = ? AND id = ?", rid, mid).Updates(upd).Error
	if err != nil {
		return nil, err
	}
	return &mitem, nil
}
