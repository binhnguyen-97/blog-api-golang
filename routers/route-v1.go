package routers

import (
	"blog-api-golang/controllers"
	"blog-api-golang/middlewares"
	"blog-api-golang/utils"

	"github.com/gin-gonic/gin"
)

func InitV1Router(r *gin.Engine) *gin.Engine {
	v1Group := r.Group("/v1")

	privateRoute := v1Group.Group("private")

	privateRoute.POST("/article", middlewares.AuthMiddleware(utils.ACCEPT_ALL_ROLES), controllers.PostArticleHandler)
	privateRoute.PUT("/article/:id", middlewares.AuthMiddleware(utils.ACCEPT_ALL_ROLES), controllers.PutArticleHandler)
	privateRoute.DELETE("/article/:id", middlewares.AuthMiddleware(utils.ACCEPT_MAINTAINER_ROLES), controllers.DeleteArticleHandler)

	privateRoute.POST("/user/create", middlewares.AuthMiddleware(utils.ACCEPT_ADMIN_ROLES), controllers.CreateAccountHandler)

	publicRoute := v1Group.Group("public")

	publicRoute.GET("/articles", controllers.GetAllArticlesHandler)
	publicRoute.GET("/article/:id", controllers.GetArticleDetailHandler)

	publicRoute.POST("/user", controllers.SignInHandler)

	return r
}
