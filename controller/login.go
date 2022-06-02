package controller

import (
	"github.com/Usigned/douyin/entity"
	"github.com/Usigned/douyin/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin
var usersLoginInfo = map[string]entity.User{
	"zhangleidouyin": {
		Id:            1,
		Name:          "zhanglei",
		FollowCount:   10,
		FollowerCount: 5,
		IsFollow:      true,
	},
	"qingcdma1330": {
		Id:            1,
		Name:          "Qing",
		FollowCount:   100,
		FollowerCount: 5000,
		IsFollow:      false,
	},
}

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
	userId, token, err := service.NewUserServiceInstance().Login(username, password)
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
	userService := service.NewUserServiceInstance()
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
