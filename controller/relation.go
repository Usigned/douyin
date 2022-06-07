package controller

import (
	"douyin/dao"
	"douyin/entity"
	"douyin/service"
	"douyin/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var relationService = service.NewRelationServiceInstance()

type UserListResponse struct {
	entity.Response
	UserList []entity.User `json:"user_list"`
}

// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {
	c.JSON(http.StatusOK, RelationActionFunc(
		c.Query("user_id"),
		c.Query("token"),
		c.Query("to_user_id"),
		c.Query("action_type"),
	))
}

func RelationActionFunc(userId, token, toUserId, actionType string) UserListResponse {
	// TODO 使用token进行鉴权
	if token == "" {
		return ErrorRelationActionResponse(utils.Error{Msg: "empty token or user_id"})
	}
	// 用户不能关注自己
	var uid int64
	var err error
	if userId == "" {
		uid, err = dao.NewLoginStatusDaoInstance().QueryUserIdByToken(token)
		if err != nil {
			return ErrorRelationActionResponse(err)
		}
	} else {
		uid, err = strconv.ParseInt(userId, 10, 64)
		if err != nil {
			return ErrorRelationActionResponse(err)
		}
	}
	tUid, err := strconv.ParseInt(toUserId, 10, 64)
	if err != nil {
		return ErrorRelationActionResponse(err)
	}

	if uid == tUid {
		return ErrorRelationActionResponse(utils.Error{
			Msg: "自己不能关注自己！",
		})
	}

	if actionType == "1" {
		err = relationService.Follow(uid, tUid, token)
		if err != nil {
			return ErrorRelationActionResponse(err)
		}
		return UserListResponse{
			Response: entity.Response{
				StatusCode: 0,
				StatusMsg:  "",
			},
		}
	} else if actionType == "2" {
		err = relationService.WithdrawFollow(uid, tUid, token)
		if err != nil {
			return ErrorRelationActionResponse(err)
		}
		return UserListResponse{
			Response: entity.Response{
				StatusCode: 0,
				StatusMsg:  "",
			},
		}
	} else {
		return UserListResponse{
			Response: entity.Response{
				StatusCode: 1,
				StatusMsg:  "Service Wrong!",
			},
		}
	}
}

func ErrorRelationActionResponse(err error) UserListResponse {
	return UserListResponse{
		Response: entity.Response{
			StatusCode: -1,
			StatusMsg:  err.Error(),
		},
		UserList: nil,
	}
}
