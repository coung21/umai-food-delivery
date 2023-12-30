package repository

import (
	"common"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *menuRepoMongo) DeleteMenuItem(ctx context.Context, id string) (int, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return 0, common.BadQueryParams
	}
	filter := bson.M{"_id": objectID}
	result, err := r.dbc.DeleteOne(ctx, filter)
	if err != nil {
		return int(result.DeletedCount), err
	}
	return int(result.DeletedCount), nil
}
