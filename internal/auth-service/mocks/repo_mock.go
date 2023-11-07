package mocks

import (
	"context"
	auth "umai-auth-service/interfaces"
	"umai-auth-service/model"
)

type RepoMock struct {
	auth.Repository
	MockFindUserByEmail  func(ctx context.Context, email string) (*model.User, error)
	MockInsertUser       func(ctx context.Context, user *model.User) (*model.User, error)
	MockFindUserByID     func(ctx context.Context, id int) (*model.User, error)
	MockUpdateRole       func(ctx context.Context, user *model.User) error
	MockInsertRestaurant func(ctx context.Context, res *model.Restaurant) (*model.Restaurant, error)
}

func (m *RepoMock) FindUserByEmail(ctx context.Context, email string) (*model.User, error) {
	return m.MockFindUserByEmail(ctx, email)
}

func (m *RepoMock) InsertUser(ctx context.Context, user *model.User) (*model.User, error) {
	return m.MockInsertUser(ctx, user)
}

func (m *RepoMock) FindUserByID(ctx context.Context, id int) (*model.User, error) {
	return m.MockFindUserByID(ctx, id)
}

func (m *RepoMock) UpdateRole(ctx context.Context, user *model.User) error {
	return m.MockUpdateRole(ctx, user)
}

func (m *RepoMock) InsertRestaurant(ctx context.Context, res *model.Restaurant) (*model.Restaurant, error) {
	return m.MockInsertRestaurant(ctx, res)
}
