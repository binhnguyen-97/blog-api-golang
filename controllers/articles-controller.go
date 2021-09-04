package controllers

import (
	"blog-api-golang/models"
	"blog-api-golang/types"
	"blog-api-golang/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllArticlesHandler(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "0"))

	if err != nil {
		page = 0
	}

	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))

	if err != nil {
		limit = 20
	}

	data, err := models.GetArticles(int64(limit), page)

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.GetSuccessMessage(err.Error()))
	} else {
		c.JSON(http.StatusOK, types.ListArticleResp{Status: "success", Data: data})
	}
}

func GetArticleDetailHandler(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, "Place  holder")
}

func PostArticleHandler(c *gin.Context) {
	var articleRequest types.InsertArticleRequest

	if c.Bind(&articleRequest) != nil {
		c.JSON(http.StatusBadRequest, utils.GetSuccessMessage("Invaid request body"))
	}

	result, err := models.AddNewArticle(articleRequest.Title, articleRequest.Description, articleRequest.Author)

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.GetSuccessMessage(err.Error()))
	}

	c.JSON(http.StatusCreated, result)
}

func PutArticleHandler(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, "Place  holder")
}

func DeleteArticleHandler(c *gin.Context) {
	articleId := c.Param("id")

	err := models.DeleteArticle(articleId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.GetErrorMessage(err.Error()))
	} else {
		c.JSON(http.StatusAccepted, utils.GetSuccessMessage("Deleted successfully"))
	}
}
