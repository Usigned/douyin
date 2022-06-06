package controller

import (
	"douyin/dao"
	"douyin/entity"
	"douyin/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func FollowList(c *gin.Context) {
	c.JSON(http.StatusOK, FollowListFunc(
		c.Query("user_id"),
		c.Query("token"),
	))
}

func FollowerList(c *gin.Context) {
	c.JSON(http.StatusOK, FollowerListFunc(
		c.Query("user_id"),
		c.Query("token"),
	))
}

func FollowListFunc(userId, token string) UserListResponse {
	// TODO 使用token进行鉴权
	if token == "" {
		return ErrorRelationActionResponse(utils.Error{Msg: "empty token or user_id"})
	}
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
	list, err := relationService.FollowList(uid, token)
	if err != nil {
		return ErrorFollowListResponse(err)
	}
	return UserListResponse{
		Response: entity.Response{
			StatusCode: 0,
			StatusMsg:  "Pull Follow Success!",
		},
		UserList: list,
	}
}

func FollowerListFunc(userId, token string) UserListResponse {
	// TODO 使用token进行鉴权
	if token == "" {
		return ErrorFollowerListResponse(utils.Error{Msg: "empty token or user_id"})
	}
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
	list, err := relationService.FollowerList(uid, token)
	if err != nil {
		return ErrorFollowerListResponse(err)
	}
	return UserListResponse{
		Response: entity.Response{
			StatusCode: 0,
			StatusMsg:  "Pull Follower Success!",
		},
		UserList: list,
	}
}

func ErrorFollowListResponse(err error) UserListResponse {
	return UserListResponse{
		Response: entity.Response{
			StatusCode: -1,
			StatusMsg:  err.Error(),
		},
		UserList: nil,
	}
}

func ErrorFollowerListResponse(err error) UserListResponse {
	return UserListResponse{
		Response: entity.Response{
			StatusCode: -1,
			StatusMsg:  err.Error(),
		},
		UserList: nil,
	}
}
