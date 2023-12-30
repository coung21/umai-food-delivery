package repository

import (
	"context"
	"menu-service/model"
)

func (r *menuRepoMongo) InsertMenuItem(ctx context.Context, mitem *model.MenuItem) (interface{}, error) {
	result, err := r.dbc.InsertOne(ctx, mitem)
	if err != nil {
		return nil, err
	}
	return result.InsertedID, nil
}
