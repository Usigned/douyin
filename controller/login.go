package controller

import (
	"douyin/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserLoginResponse struct {
	entity.Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	c.JSON(http.StatusOK, LoginFunc(username, password))
}

// Register /douyin/user/register/
func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	c.JSON(http.StatusOK, RegisterFunc(username, password))
}

func LoginFunc(username, password string) UserLoginResponse {
	userId, token, err := userService.Login(username, password)
	if err != nil {
		return ErrorUserLoginResponse(err)
	}
	return UserLoginResponse{
		Response: entity.Response{
			StatusCode: 0,
			StatusMsg:  "success",
		},
		UserId: *userId,
		Token:  *token,
	}
}

func RegisterFunc(username, password string) UserLoginResponse {
	if err := userService.Register(username, password); err != nil {
		return ErrorUserLoginResponse(err)
	}
	return LoginFunc(username, password)
}

func ErrorUserLoginResponse(err error) UserLoginResponse {
	return UserLoginResponse{
		Response: entity.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		},
	}
}
