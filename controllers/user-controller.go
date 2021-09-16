package controllers

import (
	"blog-api-golang/cache"
	"blog-api-golang/models"
	"blog-api-golang/services"
	"blog-api-golang/types"
	"blog-api-golang/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func SignInHandler(c *gin.Context) {
	var signInRequest types.SignInRequest

	if c.Bind(&signInRequest) != nil {
		c.JSON(http.StatusBadRequest, utils.GetErrorMessage("Invalid Params"))
	}

	userInfo, err := models.GetUserInfo(signInRequest.Email)

	if err != nil {
		c.JSON(http.StatusForbidden, utils.GetErrorMessage("Username is not existed"))
		return
	}

	check := utils.CheckPasswordHash(signInRequest.Password, userInfo.Password)

	if !check {
		c.JSON(http.StatusForbidden, utils.GetErrorMessage("Username/ Password is not correct"))
		return
	}

	jwtToken, err := utils.GenerateJWT(userInfo.Email, userInfo.Role)

	if err != nil {
		c.JSON(http.StatusForbidden, utils.GetErrorMessage("Username/ Password is not correct"))
		return
	}

	c.JSON(http.StatusOK, utils.GetSuccessMessage(types.AuthenticateResp{
		Email:        userInfo.Email,
		PrivateToken: jwtToken,
		Role:         userInfo.Role,
	}))
}

func CreateAccountHandler(c *gin.Context) {
	var createAccountReq types.CreateAccountRequest

	if c.Bind(&createAccountReq) != nil {
		c.JSON(http.StatusBadRequest, utils.GetErrorMessage("Invalid Params"))
	}

	if createAccountReq.Role == "" {
		createAccountReq.Role = utils.USER
	}

	userInfo, _ := models.GetUserInfo(createAccountReq.Email)

	if userInfo.Email != "" {
		c.JSON(http.StatusForbidden, utils.GetErrorMessage("Email already existed"))
		return
	}

	createdUser, err := models.CreateUser(createAccountReq.Email, createAccountReq.Password, createAccountReq.Role)

	if err != nil {
		c.JSON(http.StatusForbidden, utils.GetErrorMessage("Fail to create"))
		return
	}

	mailList := []string{createdUser.Email}
	otp := utils.OTPGenerator()

	if ok := cache.SetValue(createdUser.Email, otp, time.Minute*5); !ok {
		c.JSON(http.StatusBadRequest, utils.GetErrorMessage("Error when generate OTP"))
	}

	mailService := services.CreateNewMail(mailList, "Welcome to The Bidu family")

	err = mailService.SendMail("template/register.html", types.RegisterTemplateItems{
		Email: createdUser.Email,
		OTP:   otp,
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.GetErrorMessage("Sent OTP Fail with Error "+err.Error()))
		return
	}

	c.JSON(http.StatusCreated, createdUser)
}

func GetUserInfoHandler(c *gin.Context) {

	userEmail := c.GetHeader("email")

	if userEmail == "" {
		c.JSON(http.StatusForbidden, utils.GetErrorMessage("User is not existed"))
		return
	}

	userInfo, err := models.GetUserInfo(userEmail)

	if err != nil {
		c.JSON(http.StatusForbidden, utils.GetErrorMessage("User is not existed"))
		return
	}

	c.JSON(http.StatusOK, utils.GetSuccessMessage(userInfo))
}

func GetAllUserHandler(c *gin.Context) {
	userList, err := models.GetAllUser()

	if err != nil {
		c.JSON(http.StatusForbidden, utils.GetErrorMessage(err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.GetSuccessMessage(userList))
}
