package mocks

import (
	"context"
	menu "menu-service/interfaces"
	"menu-service/model"
)

type RepoMock struct {
	menu.Repository
	MockInsertMenuItem      func(ctx context.Context, mitem *model.MenuItem) (interface{}, error)
	MockListMenuItemByResID func(ctx context.Context, rid int) ([]model.MenuItem, error)
	MockUpdateMenuItem      func(ctx context.Context, rid int, mid string, upd *model.UpdateMenuItem) (*model.MenuItem, error)
	MockDeleteMenuItem      func(ctx context.Context, id string) (int, error)
	MockFindMenuItemByID    func(ctx context.Context, id string) (*model.MenuItem, error)
}

func (m *RepoMock) InsertMenuItem(ctx context.Context, mitem *model.MenuItem) (interface{}, error) {
	return m.MockInsertMenuItem(ctx, mitem)
}

func (m *RepoMock) ListMenuItemByResID(ctx context.Context, rid int) ([]model.MenuItem, error) {
	return m.MockListMenuItemByResID(ctx, rid)
}
func (m *RepoMock) UpdateMenuItem(ctx context.Context, rid int, mid string, upd *model.UpdateMenuItem) (*model.MenuItem, error) {
	return m.UpdateMenuItem(ctx, rid, mid, upd)
}
func (m *RepoMock) DeleteMenuItem(ctx context.Context, id string) (int, error) {
	return m.DeleteMenuItem(ctx, id)
}
func (m *RepoMock) FindMenuItemByID(ctx context.Context, id string) (*model.MenuItem, error) {
	return m.MockFindMenuItemByID(ctx, id)
}
