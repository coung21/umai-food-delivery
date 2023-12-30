package repository

import "go.mongodb.org/mongo-driver/mongo"

type menuRepoMongo struct {
	dbc *mongo.Collection
}

func NewMenuRepo(dbc *mongo.Collection) *menuRepoMongo {
	return &menuRepoMongo{dbc: dbc}
}
