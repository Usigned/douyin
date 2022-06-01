package controller

import (
	"github.com/Usigned/douyin/entity"
	"github.com/Usigned/douyin/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var userIdSequence = int64(1)

type UserResponse struct {
	entity.Response
	User entity.User `json:"user"`
}

// UserInfo /douyin/user
func UserInfo(c *gin.Context) {
	c.JSON(http.StatusOK, UserInfoFunc(
		c.Query("token"),
		c.Query("user_id"),
	))
}

func UserInfoFunc(token, userId string) UserResponse {
	// TODO
	// 判断用户是否存在，存在则返回用户信息
	uid, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		return ErrorUserResponse(err)
	}

	user, err := service.NewUserServiceInstance().FindUserById(uid)
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
