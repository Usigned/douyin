package controller

import (
	"errors"
	"fmt"
	"github.com/Usigned/douyin/dao"
	"github.com/Usigned/douyin/entity"
	"github.com/Usigned/douyin/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
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
	// 用户输入验证
	err := InfoVerify(username, password)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: entity.Response{StatusCode: 1, StatusMsg: err.Error()},
		})
		return
	}
	token := "<" + username + "><" + password + ">"
	// 先查缓存 ..
	if _, exist := usersLoginInfo[token]; !exist {
		if user, _ := userService.FindUserByName(username); user == nil {
			// 用户注册
			userIdSequence, _ := userService.LastId()
			atomic.AddInt64(&userIdSequence, 1)
			newUser := &dao.User{
				Id:       userIdSequence,
				Name:     username,
				Password: password,
			}
			var err error
			user, err = userService.SaveUser(newUser)
			if err != nil {
				c.JSON(http.StatusOK, UserLoginResponse{
					Response: entity.Response{StatusCode: 1, StatusMsg: "User register failed, Please retry for a minute!"},
				})
				return
			}
			usersLoginInfo[token] = *user
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: entity.Response{StatusCode: 0},
				UserId:   userIdSequence,
				Token:    token,
			})
			return
		}
	}
	c.JSON(http.StatusOK, UserLoginResponse{
		Response: entity.Response{StatusCode: 1, StatusMsg: "User already exist, don't register again!"},
	})
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	err := InfoVerify(username, password)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: entity.Response{StatusCode: 1, StatusMsg: err.Error()},
		})
		return
	}
	token := "<" + username + "><" + password + ">"
	// 先查询缓存 ..
	if _, exist := usersLoginInfo[token]; !exist {
		user, _ := userService.FindUserByName(username)
		if user == nil {
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: entity.Response{StatusCode: 1, StatusMsg: "User doesn't exist, Please Register"},
			})
			return
		}
		usersLoginInfo[token] = *user
	}
	// 密码校验
	result, _ := userService.FindUserByToken(token)
	if result == nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: entity.Response{StatusCode: 1, StatusMsg: "Password Wrong!"},
		})
		return
	}
	c.JSON(http.StatusOK, UserLoginResponse{
		Response: entity.Response{StatusCode: 0},
		UserId:   usersLoginInfo[token].Id,
		Token:    token,
	})
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

func InfoVerify(username string, password string) error {
	if Check(username) {
		return errors.New("Please Check Username!\nThe length is controlled within 4-32 characters, and <, >, \\is not allowed")
	}
	if Check(password) {
		return errors.New("Please Check Password!\nThe length is controlled within 4-32 characters, and <, >, \\is not allowed")
	}
	return nil
}

func Check(str string) bool {
	lenth := len(str)
	if lenth < 4 || lenth > 32 {
		return true
	}
	if strings.Contains(str, "<") || strings.Contains(str, ">") ||
		strings.Contains(str, "/") || strings.Contains(str, "\\") {
		return true
	}
	return false
}
