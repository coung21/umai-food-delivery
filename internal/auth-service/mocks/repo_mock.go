package mocks

import (
	"context"
	"time"
	auth "umai-auth-service/interfaces"
	"umai-auth-service/model"
)

type RepoMock struct {
	auth.Repository
	MockFindUserByEmail    func(ctx context.Context, email string) (*model.User, error)
	MockInsertUser         func(ctx context.Context, user *model.User) (int, error)
	MockFindUserByID       func(ctx context.Context, id int) (*model.User, error)
	MockUpdateRole         func(ctx context.Context, user *model.User) error
	MockInsertRestaurant   func(ctx context.Context, res *model.Restaurant) (int, error)
	MockUpdateUser         func(ctx context.Context, olduser *model.User, upd *model.UserUpdate) (*model.User, error)
	MockFindRestaurantByID func(ctx context.Context, id int) (*model.Restaurant, error)
	MockUpdateRestaurant   func(ctx context.Context, oldres *model.Restaurant, upd *model.RestaurantUpdate) (*model.Restaurant, error)
}

func (m *RepoMock) FindUserByEmail(ctx context.Context, email string) (*model.User, error) {
	return m.MockFindUserByEmail(ctx, email)
}

func (m *RepoMock) InsertUser(ctx context.Context, user *model.User) (int, error) {
	return m.MockInsertUser(ctx, user)
}

func (m *RepoMock) FindUserByID(ctx context.Context, id int) (*model.User, error) {
	return m.MockFindUserByID(ctx, id)
}

func (m *RepoMock) UpdateRole(ctx context.Context, user *model.User) error {
	return m.MockUpdateRole(ctx, user)
}

func (m *RepoMock) InsertRestaurant(ctx context.Context, res *model.Restaurant) (int, error) {
	return m.MockInsertRestaurant(ctx, res)
}

func (m *RepoMock) UpdateUser(ctx context.Context, olduser *model.User, upd *model.UserUpdate) (*model.User, error) {
	return m.MockUpdateUser(ctx, olduser, upd)
}

func (m *RepoMock) FindRestaurantByID(ctx context.Context, id int) (*model.Restaurant, error) {
	return m.MockFindRestaurantByID(ctx, id)
}

func (m *RepoMock) UpdateRestaurant(ctx context.Context, oldres *model.Restaurant, upd *model.RestaurantUpdate) (*model.Restaurant, error) {
	return m.MockUpdateRestaurant(ctx, oldres, upd)
}

type CacheRepoMock struct {
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
