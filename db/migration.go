package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Migragion() error {

	models := mongo.IndexModel{
		Keys: bson.D{
			primitive.E{
				Key:   "email",
				Value: 1,
			},
		},
		Options: options.Index().SetUnique(true),
	}

	_, err := UserCollection().Indexes().CreateOne(context.TODO(), models)

	if err != nil {
		return err
	}

	return nil
}
