package controller

import (
	"douyin/entity"
	"douyin/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var userService = service.NewUserServiceInstance()

type UserResponse struct {
	entity.Response
	User entity.User `json:"user"`
}

// UserInfo /douyin/user
func UserInfo(c *gin.Context) {
	c.JSON(http.StatusOK, UserInfoFunc(
		c.Query("user_id"),
		c.Query("token"),
	))
}

func UserInfoFunc(userId, token string) UserResponse {
	uid, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		return ErrorUserResponse(err)
	}

	user, err := service.NewUserServiceInstance().UserInfo(uid)
	if err != nil {
		return ErrorUserResponse(err)
	}
	if user == nil {
		return FailUserResponse("user not exist: uid " + strconv.FormatInt(uid, 10))
	}
	return UserResponse{
		Response: entity.Response{
			StatusCode: 0,
			StatusMsg:  "success",
		},
		User: *user,
	}
}

func ErrorUserResponse(err error) UserResponse {
	return UserResponse{
		Response: entity.Response{
			StatusCode: -1,
			StatusMsg:  err.Error(),
		},
	}
}

func FailUserResponse(msg string) UserResponse {
	return UserResponse{
		Response: entity.Response{
			StatusCode: -1,
			StatusMsg:  msg,
		},
	}
}
