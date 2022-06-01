package controller

import (
	"github.com/Usigned/douyin/entity"
	"github.com/Usigned/douyin/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync/atomic"
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

	token := username + password

	if user, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: entity.Response{StatusCode: 0, StatusMsg: "success"},
			UserId:   user.Id,
			Token:    token,
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: entity.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}

func LoginFunc(username, password string) UserLoginResponse {
	// TODO
	return FailUserLoginResponse("wrong password")
}

// Register /douyin/user/register/
func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token := username + password

	if _, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: entity.Response{StatusCode: 1, StatusMsg: "User already exist"},
		})
	} else {
		atomic.AddInt64(&userIdSequence, 1)
		newUser := entity.User{
			Id:   userIdSequence,
			Name: username,
		}
		usersLoginInfo[token] = newUser
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: entity.Response{StatusCode: 0},
			UserId:   userIdSequence,
			Token:    username + password,
		})
	}
}

func RegisterFunc(username, password string) UserLoginResponse {
	// register
	userService := service.NewUserServiceInstance()
	user, err := userService.FindUserByName(username)
	if err != nil {
		return ErrorUserLoginResponse(err)
	}
	if user != nil {
		return FailUserLoginResponse("username already exist: " + username)
	}

	if err = userService.AddUser(username, password); err != nil {
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

func FailUserLoginResponse(msg string) UserLoginResponse {
	return UserLoginResponse{
		Response: entity.Response{
			StatusCode: -1,
			StatusMsg:  msg,
		},
	}
}
