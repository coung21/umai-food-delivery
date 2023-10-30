package mocks

import (
	"context"
	auth "umai-auth-service/interfaces"
	"umai-auth-service/model"
)

type UserRepoMock struct {
	auth.Repository
	MockFindUserByEmail func(ctx context.Context, email string) (*model.User, error)
	MockInsertUser      func(ctx context.Context, user *model.User) (*model.User, error)
}

func (m *UserRepoMock) FindUserByEmail(ctx context.Context, email string) (*model.User, error) {
	return m.MockFindUserByEmail(ctx, email)
}

func (m *UserRepoMock) InsertUser(ctx context.Context, user *model.User) (*model.User, error) {
	return m.MockInsertUser(ctx, user)
}
