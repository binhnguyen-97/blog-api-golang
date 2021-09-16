package models

import (
	"blog-api-golang/db"
	"blog-api-golang/types"
	"blog-api-golang/utils"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetArticles(limit int64, page int) ([]types.ArticleResp, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	limitState := utils.GetLimitStage(int(limit))
	lookupStage := utils.GetLookupStage("writer", "author", "_id", "author")
	unwindStage := utils.GetUnwindStage("$author", false)

	cursor, err := db.ArticleCollection().Aggregate(ctx, mongo.Pipeline{lookupStage, limitState, unwindStage})

	if err != nil {
		panic(err)
	}

	var results []types.ArticleResp

	if err = cursor.All(ctx, &results); err != nil {
		panic(err)
	}

	return results, err
}

func GetArticleDetail(id string) ([]types.ArticleResp, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	articleObjectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return []types.ArticleResp{}, err
	}

	matchCondition := bson.D{
		primitive.E{
			Key:   "_id",
			Value: articleObjectId,
		},
	}

	matchState := utils.GetMatchStage(matchCondition)
	limitState := utils.GetLimitStage(1)
	lookupStage := utils.GetLookupStage("writer", "author", "_id", "author")
	unwindStage := utils.GetUnwindStage("$author", false)

	cursor, err := db.ArticleCollection().Aggregate(ctx, mongo.Pipeline{matchState, limitState, lookupStage, unwindStage})

	if err != nil {
		return []types.ArticleResp{}, err
	}

	var results []types.ArticleResp

	if err = cursor.All(ctx, &results); err != nil {
		return []types.ArticleResp{}, err
	}

	return results, nil
}

func AddNewArticle(title string, shortDescription string, author string, content string) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	articleCollection := db.ArticleCollection()
	authorId, err := primitive.ObjectIDFromHex(author)

	if err != nil {
		return nil, err
	}

	article := types.Article{
		Title:            title,
		ShortDescription: shortDescription,
		Content:          content,
		Author:           authorId,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
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

func UpdateArticle(id string, title string, shortDescription string, author string, content string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	articleId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	authorId, err := primitive.ObjectIDFromHex(author)

	if err != nil {
		return err
	}

	filter := bson.M{"_id": articleId}
	update := bson.D{
		primitive.E{
			Key: "$set",
			Value: bson.M{
				"title":       title,
				"description": shortDescription,
				"content":     content,
				"author":      authorId,
				"updatedAt":   time.Now(),
			},
		},
	}

	_, err = db.ArticleCollection().UpdateOne(
		ctx,
		filter,
		update,
	)

	if err != nil {
		return err
	}
	return nil
}
