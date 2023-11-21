package api

import (
	"github.com/gin-gonic/gin"
	"xapi/model"
	"xapi/service"
	"xapi/utils"
)

func SetupUserRouter(router *gin.Engine) {
	v1 := router.Group("/v1/user")
	v1.POST("login", login)
	v1.POST("register", register)
	v1.POST("info", info)
}

func login(c *gin.Context) {

}

// @title Your Project API
// @version 1.0
// @description This is a sample service for Your Project.

// @Summary User register
// @Description Logs the user in and returns an access token.
// @ID register1
// @Accept  json
// @Produce  json
// @Param username query string true "Username"
// @Param password query string true "Password"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Bad Request"
// @Router /register [post]
func register(c *gin.Context) {
	var user model.User
	if err := c.ShouldBind(&user); err != nil {
		ErrorResp(c, 405, "参数缺失")
	}
	user.Username = "fchpy"
	user.Password = "123456"

	err := service.GetDB().Create(&user).Error
	if err != nil {

	}
	token, err := utils.GetToken(user.ID)
	if err != nil {

	}

	SuccessResp(c, "注册成功", gin.H{"token": token})
}

func info(c *gin.Context) {

}
