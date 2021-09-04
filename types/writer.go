package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type Writer struct {
	Id     primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name   string             `bson:"name" json:"name"`
	Avatar string             `bson:"avatar,omitempty" json:"avatar,omitempty"`
}
