package mocks

import (
	"context"
	menu "menu-service/interfaces"
	"menu-service/model"
	"time"
)

type CacheRepoMock struct {
	menu.CacheRepository
	MockSet func(ctx context.Context, id string, res *model.MenuItem, ttl time.Duration) error
	MockGet func(ctx context.Context, id string) (*model.MenuItem, error)
	MockDel func(ctx context.Context, id string) error
}

func (m *CacheRepoMock) Set(ctx context.Context, id string, res *model.MenuItem, ttl time.Duration) error {
	return m.MockSet(ctx, id, res, time.Hour)
}

func (m *CacheRepoMock) Get(ctx context.Context, id string) (*model.MenuItem, error) {
	return m.MockGet(ctx, id)
}

func (m *CacheRepoMock) Del(ctx context.Context, id string) error {
	return m.MockDel(ctx, id)
}
