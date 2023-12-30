package repository

import (
	"common"
	"context"
	"menu-service/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *menuRepoMongo) UpdateMenuItem(ctx context.Context, rid int, mid string, upd *model.UpdateMenuItem) (*model.MenuItem, error) {
	var newData model.MenuItem
	objectID, err := primitive.ObjectIDFromHex(mid)
	if err != nil {
		return nil, common.BadQueryParams
	}
	filter := bson.M{"_id": objectID, "restaurant_id": rid}
	options := options.FindOneAndUpdate().SetReturnDocument(options.After)
	err = r.dbc.FindOneAndUpdate(ctx, filter, bson.M{"$set": upd}, options).Decode(&newData)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, common.NotFound
		}
		return nil, err
	}

	return &newData, nil
}
