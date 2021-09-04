package routers

import (
	"blog-api-golang/controllers"

	"github.com/gin-gonic/gin"
)

func InitV1Router(r *gin.Engine) *gin.Engine {
	v1Group := r.Group("/v1")

	privateRoute := v1Group.Group("private")

	privateRoute.PUT("/article/:id", controllers.PutArticleHandler)
	privateRoute.DELETE("/article/:id", controllers.DeleteArticleHandler)
	privateRoute.POST("/article", controllers.PostArticleHandler)

	publicRoute := v1Group.Group("public")

	publicRoute.GET("/articles", controllers.GetAllArticlesHandler)
	publicRoute.GET("/article/:id", controllers.GetArticleDetailHandler)

	return r
}
