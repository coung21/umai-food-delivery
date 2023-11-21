package repository

import "go.mongodb.org/mongo-driver/mongo"

type menuRepo struct {
	dbc *mongo.Collection
}

func NewMenuRepo(dbc *mongo.Collection) *menuRepo {
	return &menuRepo{dbc: dbc}
}
