package routers

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/heath-check", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})

	InitV1Router(r)

	return r
}
