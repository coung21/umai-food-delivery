package repository

import (
	"common"
	"context"
	"menu-service/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *menuRepo) FindMenuItemByID(ctx context.Context, id string) (*model.MenuItem, error) {
	var mitem model.MenuItem
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, common.BadQueryParams
	}

	err = r.dbc.FindOne(ctx, bson.M{"_id": objectID}).Decode(&mitem)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, common.NotFound
		}
		return nil, err
	}
	return &mitem, nil
}
