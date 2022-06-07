package controller

import (
	"douyin/entity"
	"douyin/pack"
	"douyin/service"
	"douyin/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type VideoListResponse struct {
	entity.Response
	VideoList []entity.Video `json:"video_list"`
}

// PublishList /douyin/publish
func PublishList(c *gin.Context) {
	c.JSON(http.StatusOK, publishListFunc(c.Query("token"), c.Query("user_id")))
}

func publishListFunc(token, userId string) VideoListResponse {
	// TODO 使用token进行鉴权
	if token == "" || userId == "" {
		return ErrorVideoListResponse(utils.Error{Msg: "empty token or user_id"})
	}
	uid, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		return ErrorVideoListResponse(err)
	}

	videos, err := service.NewVideoServiceInstance().PublishList(uid)
	if err != nil {
		return ErrorVideoListResponse(err)
	}

	return VideoListResponse{
		Response: entity.Response{
			StatusCode: 0,
			StatusMsg:  "success",
		},
		VideoList: pack.VideoPtrs(videos),
	}

}

func ErrorVideoListResponse(err error) VideoListResponse {
	return VideoListResponse{
		Response: entity.Response{
			StatusCode: -1,
			StatusMsg:  err.Error(),
		},
		VideoList: nil,
	}
}
