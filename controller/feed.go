package controller

import (
	"douyin/entity"
	"douyin/pack"
	"douyin/service"
	"douyin/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type FeedResponse struct {
	entity.Response
	VideoList []entity.Video `json:"video_list"`
	NextTime  int64          `json:"next_time"`
}

// Feed use userService and videoService to query data
func Feed(c *gin.Context) {
	c.JSON(http.StatusOK, FeedFunc(c.Query("latest_time"), c.Query("token")))
}

func FeedFunc(latestTime string, token string) FeedResponse {
	timeInt, _ := strconv.ParseInt(latestTime, 10, 64)
	nextTime, videos, err := service.NewVideoServiceInstance().Feed(timeInt, token, utils.DefaultLimit)
	// service层出错
	if err != nil {
		return ErrorFeedResponse(err)
	}

	return FeedResponse{
		Response: entity.Response{
			StatusCode: 0,
			StatusMsg:  "success",
		},
		VideoList: pack.VideoPtrs(videos),
		NextTime:  *nextTime,
	}
}

func ErrorFeedResponse(err error) FeedResponse {
	return FeedResponse{
		Response: entity.Response{
			StatusCode: -1,
			StatusMsg:  err.Error(),
		},
		VideoList: nil,
		NextTime:  time.Now().UnixMilli(),
	}
}
