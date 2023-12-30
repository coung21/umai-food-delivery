package repository

import (
	"common"
	"context"
	"menu-service/model"

	"gorm.io/gorm"
)

func (r *menuRepoMysql) FindMenuItemByID(ctx context.Context, id int) (*model.MenuItem, error) {
	var mitem model.MenuItem
	err := r.db.Where("id = ?", id).First(&mitem).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.NotFound
		}
		return nil, err
	}
	return &mitem, nil
}
