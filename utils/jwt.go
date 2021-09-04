package utils

import (
	"blog-api-golang/config"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateJWT(email, role string) (string, error) {
	var mySigningKey = []byte(config.Config.JWT_SECRET_KEY)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["email"] = email
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		log.Fatalf("Something Went Wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}
