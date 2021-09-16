package types

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ArticleResp struct {
	ID               primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title            string             `bson:"title" json:"title"`
	Author           Writer             `bson:"author" json:"author"`
	ShortDescription string             `bson:"short_description" json:"shortDescription"`
	Content          string             `bson:"content" json:"content"`
	CreatedAt        time.Time          `bson:"createdAt" json:"createdAt"`
	UpdateAt         time.Time          `bson:"updatedAt" json:"updatedAt"`
}

type ListArticleResp struct {
	Status string        `json:"status"`
	Data   []ArticleResp `json:"data"`
}

type AuthenticateResp struct {
	Email        string `json:"email"`
	Name         string `json:"name"`
	PrivateToken string `json:"privateToken"`
	Role         string `json:"role"`
}

type ErrorRespone struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type SuccessRespone struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}
