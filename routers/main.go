package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.New()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(cors.New(config))

	router.GET("/heath-check", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})

	InitV1Router(router)

	return router
}
