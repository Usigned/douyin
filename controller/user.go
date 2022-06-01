package controller

import (
	"fmt"
	"github.com/Usigned/douyin/dao"
	"github.com/Usigned/douyin/entity"
	"github.com/Usigned/douyin/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
}
var curUser dao.User
var userService = service.NewUserServiceInstance()

type UserLoginResponse struct {
	entity.Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	entity.Response
	User entity.User `json:"user"`
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	token := username + password
	// 先查缓存 ..
	if _, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: entity.Response{StatusCode: 1, StatusMsg: "User already exist, don't register again!"},
		})
	} else {
		userIdSequence, _ := userService.LastId()
		fmt.Println("id is:", userIdSequence)
		atomic.AddInt64(&userIdSequence, 1)
		newUser := entity.User{
			Id:   userIdSequence,
			Name: username,
		}
		curUser.Id = newUser.Id
		curUser.Name = username
		curUser.Password = password
		err := userService.SaveUser(&curUser)
		if err != nil {
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: entity.Response{StatusCode: 1, StatusMsg: "User already exist, don't register again!"},
			})
		}
		usersLoginInfo[token] = newUser
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: entity.Response{StatusCode: 0},
			UserId:   userIdSequence,
			Token:    username + password,
		})
	}
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token := username + password
	// 先查询缓存 ..
	if user, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: entity.Response{StatusCode: 0},
			UserId:   user.Id,
			Token:    token,
		})
	} else {
		result, err := userService.FindUserByName(username)
		if err != nil {
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: entity.Response{StatusCode: 1, StatusMsg: "User doesn't exist, Please Register"},
			})
		} else {
			usersLoginInfo[token] = *result
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: entity.Response{StatusCode: 0},
				UserId:   result.Id,
				Token:    token,
			})
		}
	}

}

func UserInfo(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	token := c.Query("token")
	// 先查询缓存 ..
	if user, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, UserResponse{
			Response: entity.Response{StatusCode: 0},
			User:     user,
		})
	} else {
		// 查询数据库
		result, err := userService.FindUserById(id)
		if err != nil {
			c.JSON(http.StatusOK, UserResponse{
				Response: entity.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
			})
		} else {
			usersLoginInfo[token] = *result
			c.JSON(http.StatusOK, UserResponse{
				Response: entity.Response{StatusCode: 0},
				User:     *result,
			})
		}
	}
}
