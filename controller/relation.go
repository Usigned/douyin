package controller

import (
	"github.com/Usigned/douyin/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserListResponse struct {
	entity.Response
	UserList []entity.User `json:"user_list"`
}

// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {
	token := c.Query("token")

	if _, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, entity.Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, entity.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// FollowList all users have same follow list
func FollowList(c *gin.Context) {
	c.JSON(http.StatusOK, UserListResponse{
		Response: entity.Response{
			StatusCode: 0,
		},
		UserList: []entity.User{DemoUser},
	})
}

// FollowerList all users have same follower list
func FollowerList(c *gin.Context) {
	c.JSON(http.StatusOK, UserListResponse{
		Response: entity.Response{
			StatusCode: 0,
		},
		UserList: []entity.User{DemoUser},
	})
}
