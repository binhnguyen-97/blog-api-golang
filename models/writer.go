package models

import (
	"blog-api-golang/db"
	"blog-api-golang/types"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func GetAllWriter() ([]types.Writer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	cursor, err := db.WritterCollection().Find(ctx, bson.D{})

	if err != nil {
		panic(err)
	}

	var results []types.Writer

	if err = cursor.All(ctx, &results); err != nil {
		panic(err)
	}

	return results, err
}
