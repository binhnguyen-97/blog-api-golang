package types

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ArticleResp struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title       string             `bson:"title" json:"title"`
	Author      Writer             `bson:"author" json:"author"`
	Description string             `bson:"description" json:"description"`
	CreatedAt   time.Time          `bson:"createdAt" json:"createdAt"`
}

type ListArticleResp struct {
	Status string        `json:"status"`
	Data   []ArticleResp `json:"data"`
}

type ErrorRespone struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type SuccessRespone struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}
