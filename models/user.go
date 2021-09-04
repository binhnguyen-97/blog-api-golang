package models

import (
	"blog-api-golang/db"
	"blog-api-golang/types"
	"blog-api-golang/utils"
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func CreateUser(email string, password string, role string) (types.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	userInfo := types.User{
		Email:     email,
		Role:      role,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	hashedPassword, err := utils.GeneratehashPassword(password)

	if err != nil {
		return types.User{}, err
	}

	userInfo.Password = hashedPassword

	insertResult, err := db.UserCollection().InsertOne(ctx, userInfo)

	if err != nil || insertResult.InsertedID == "" {
		return types.User{}, err
	}

	return userInfo, nil
}

func GetUserInfo(email string) (types.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	userFilter := bson.M{
		"email": email,
	}

	var userInfo types.User

	db.UserCollection().FindOne(ctx, userFilter).Decode(&userInfo)

	if userInfo.Email == "" {
		return types.User{}, errors.New("user is not existed")
	}

	return userInfo, nil
}
