package controller

import (
	"douyin/dao"
	"douyin/entity"
	"douyin/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AttentionController struct{}

type UserListResponse struct {
	entity.Response
	UserList []dao.UserList `json:"user_list"`
}

func FollowList(c *gin.Context) {
	query := dao.Query{}
	err := c.BindQuery(&query)
	if err != nil {
		return
	}
	var list []dao.UserList
	list, _ = service.FindFollowList(query)
	//fmt.Println(list)
	c.JSON(
		http.StatusOK,
		UserListResponse{
			Response: entity.Response{StatusCode: 0, StatusMsg: "查询成功"},
			UserList: list,
		},
	)
}

func FollowerList(c *gin.Context) {
	query := dao.Query{}
	err := c.BindQuery(&query)
	if err != nil {
		return
	}
	var list []dao.UserList
	list, _ = service.FindFollowerList(query)
	//fmt.Println(list)
	c.JSON(
		http.StatusOK,
		UserListResponse{
			Response: entity.Response{StatusCode: 0, StatusMsg: "查询成功"},
			UserList: list,
		},
	)
}

// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {
	change := dao.Change{}
	err := c.BindQuery(&change)
	if err != nil {
		return
	}
	if _, exist := usersLoginInfo[change.Token]; exist {
		err := service.RelationAction(change)
		if err != nil {
			return
		} else {
			c.JSON(http.StatusOK, entity.Response{StatusCode: 0})
		}
	} else {
		c.JSON(http.StatusOK, entity.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}
