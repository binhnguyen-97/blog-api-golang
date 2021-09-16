package controllers

import (
	"blog-api-golang/models"
	"blog-api-golang/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllWriterHandler(c *gin.Context) {
	data, err := models.GetAllWriter()

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.GetSuccessMessage(err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.GetSuccessMessage(data))
}
