package models

import (
	"blog-api-golang/db"
	"blog-api-golang/types"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetArticles(limit int64, page int) ([]types.ArticleResp, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	lookupStage := bson.D{
		primitive.E{
			Key: "$lookup",
			Value: bson.D{
				primitive.E{
					Key:   "from",
					Value: "writer",
				},
				primitive.E{
					Key:   "localField",
					Value: "author",
				},
				primitive.E{
					Key:   "foreignField",
					Value: "_id",
				},
				primitive.E{
					Key:   "as",
					Value: "author",
				},
			},
		},
	}

	unwindStage := bson.D{
		primitive.E{
			Key: "$unwind",
			Value: bson.D{
				primitive.E{
					Key:   "path",
					Value: "$author",
				},
				primitive.E{
					Key:   "preserveNullAndEmptyArrays",
					Value: false,
				},
			},
		},
	}

	limitStage := bson.D{
		primitive.E{
			Key:   "$limit",
			Value: limit,
		},
	}

	cursor, err := db.ArticleCollection().Aggregate(ctx, mongo.Pipeline{lookupStage, limitStage, unwindStage})

	if err != nil {
		panic(err)
	}

	var results []types.ArticleResp

	if err = cursor.All(ctx, &results); err != nil {
		panic(err)
	}

	return results, err
}

func AddNewArticle(title string, description string, author string) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	articleCollection := db.ArticleCollection()
	authorId, err := primitive.ObjectIDFromHex(author)

	if err != nil {
		return nil, err
	}

	article := types.Article{
		Title:       title,
		Description: description,
		Author:      authorId,
		CreatedAt:   time.Now(),
	}

	insertResult, err := articleCollection.InsertOne(ctx, article)

	if err != nil {
		return nil, err
	}

	return insertResult, nil
}

func DeleteArticle(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	articleObjectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	_, err = db.
		ArticleCollection().
		DeleteOne(ctx,
			bson.D{
				primitive.E{
					Key:   "_id",
					Value: articleObjectId,
				},
			},
		)
	if err != nil {
		return err
	}
	return nil
}
