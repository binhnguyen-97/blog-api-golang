package routers

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "https://blog-admin.thebidufamily.com", "https://blog.thebidufamily.com", "https://blog-admin-ui.vercel.app/"},
		AllowMethods:     []string{"PUT", "GET", "PATH", "OPTIONS", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/heath-check", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})

	InitV1Router(router)

	return router
}
