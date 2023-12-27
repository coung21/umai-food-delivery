package repository

import (
	"context"
	"menu-service/model"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *menuRepo) ListMenuItemByResID(ctx context.Context, rid int) ([]model.MenuItem, error) {
	var menuitems []model.MenuItem
	result, err := r.dbc.Find(ctx, bson.M{"restaurant_id": rid})
	if err != nil {
		return nil, err
	}
	defer result.Close(ctx)

	if err := result.All(ctx, &menuitems); err != nil {
		// if err == mongo.ErrNoDocuments {
		// 	return menuitems, nil
		// }
		return nil, err
	}
	if len(menuitems) == 0 {
		return make([]model.MenuItem, 0), nil
	}
	return menuitems, nil
}

func (r *menuRepo) ListMenuItemByCategory(ctx context.Context, category string) ([]model.MenuItem, error) {
	var menuitems []model.MenuItem
	result, err := r.dbc.Find(ctx, bson.M{"category": category})
	if err != nil {
		return nil, err
	}
	defer result.Close(ctx)

	if err := result.All(ctx, &menuitems); err != nil {
		return nil, err
	}
	if len(menuitems) == 0 {
		return make([]model.MenuItem, 0), nil
	}
	return menuitems, nil
}
