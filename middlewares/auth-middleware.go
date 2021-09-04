package middlewares

import (
	"blog-api-golang/config"
	"blog-api-golang/utils"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthMiddleware(rolesAccepted []string) gin.HandlerFunc {
	return func(c *gin.Context) {

		// Authentication Section
		jwtToken := strings.Split(c.GetHeader("Authorization"), " ")

		if len(jwtToken) <= 1 {
			c.JSON(http.StatusUnauthorized, utils.GetErrorMessage("Not Authorized"))
			c.Abort()
			return
		}

		var mySigningKey = []byte(config.Config.JWT_SECRET_KEY)
		token, err := jwt.Parse(jwtToken[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("there was an error in parsing token")
			}
			return mySigningKey, nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, utils.GetErrorMessage("Not Authorized"))
			c.Abort()
			return
		}

		// Authorization Section

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			requestRole := fmt.Sprintf("%v", claims["role"])

			hasPermission := false

			for _, role := range rolesAccepted {
				if requestRole == role {
					hasPermission = true
					break
				}
			}

			if hasPermission {
				c.Next()
				return
			}

			c.JSON(http.StatusForbidden, utils.GetErrorMessage("You don't have permission to do this action"))
			c.Abort()
		}

		c.Next()
	}
}
