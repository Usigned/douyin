package controller

import (
	"douyin/entity"
	"douyin/pack"
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
	// userId 当前用户
	uid, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		return ErrorRelationActionResponse(err)
	}
	relations, err := relationService.FollowList(uid, token)
	if err != nil {
		return ErrorFollowListResponse(err)
	}
	return UserListResponse{
		Response: entity.Response{
			StatusCode: 0,
			StatusMsg:  "Pull Follow Success!",
		},
		UserList: pack.RelationsPtrs(relations),
	}
}

func FollowerListFunc(userId, token string) UserListResponse {
	uid, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		return ErrorRelationActionResponse(err)
	}
	relations, err := relationService.FollowerList(uid, token)
	if err != nil {
		return ErrorFollowerListResponse(err)
	}
	return UserListResponse{
		Response: entity.Response{
			StatusCode: 0,
			StatusMsg:  "Pull Follower Success!",
		},
		UserList: pack.RelationsPtrs(relations),
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
