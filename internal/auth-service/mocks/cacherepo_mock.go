package mocks

import (
	"context"
	"time"
	auth "umai-auth-service/interfaces"
	"umai-auth-service/model"
)

type CacheRepoMock struct {
	auth.CacheRepository
	MockSet func(ctx context.Context, id int, res *model.Restaurant, ttl time.Duration) error
	MockGet func(ctx context.Context, id int) (*model.Restaurant, error)
	MockDel func(ctx context.Context, id int) error
}

func (m *CacheRepoMock) Set(ctx context.Context, id int, res *model.Restaurant, ttl time.Duration) error {
	return m.MockSet(ctx, id, res, time.Hour)
}

func (m *CacheRepoMock) Get(ctx context.Context, id int) (*model.Restaurant, error) {
	return m.MockGet(ctx, id)
}

func (m *CacheRepoMock) Del(ctx context.Context, id int) error {
	return m.MockDel(ctx, id)
}
