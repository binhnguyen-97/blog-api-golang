package db

import (
	"blog-api-golang/config"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var DB *mongo.Client

func ConnectToDatabase(ctx context.Context) {
	var err error

	DB, err = mongo.Connect(ctx, options.Client().ApplyURI(config.Config.MongoDb.URL))

	if err != nil {
		panic(err)
	}

	// Ping the primary
	if err := DB.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	log.Println("Connected to MongoDB")
}

func Disconnect(ctx context.Context) {
	log.Fatal("Disconnected")
	if err := DB.Disconnect(ctx); err != nil {
		panic(err)
	}
}

func blogBatabase() *mongo.Database {
	return DB.Database(config.Config.MongoDb.Database)
}

func ArticleCollection() *mongo.Collection {
	return blogBatabase().Collection(config.Config.Collections.Article)
}

func WritterCollection() *mongo.Collection {
	return blogBatabase().Collection(config.Config.Collections.Writer)
}

func UserCollection() *mongo.Collection {
	return blogBatabase().Collection(config.Config.Collections.Uer)
}
