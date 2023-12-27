package usecase

import (
	"common"
	"context"
)

func (u *menuUC) DeleteFavorite(ctx context.Context, uid int, mid string) (*string, error) {
	// 1. check if menu is exist in favorite
	if ex, _ := u.cacheRepo.GetFavorite(ctx, uid, mid); !ex {
		return nil, common.NotFound
	}
	// 2. delete favorite
	if err := u.cacheRepo.DelFavorite(ctx, uid, mid); err != nil {
		return nil, err
	}
	// 3. return favorite id
	return &mid, nil
}
