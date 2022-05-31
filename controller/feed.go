package controller

import (
	"github.com/Usigned/douyin/entity"
	"github.com/Usigned/douyin/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var (
	userService  = service.NewUserServiceInstance()
	videoService = service.NewVideoServiceInstance()
)

type FeedResponse struct {
	entity.Response
	VideoList []entity.Video `json:"video_list"`
	NextTime  int64          `json:"next_time"`
}

// Feed use userService and videoService to query data
func Feed(c *gin.Context) {

	c.JSON(http.StatusOK, FeedResponse{
		Response: entity.Response{
			StatusCode: 0,
			StatusMsg:  "success",
		},
		VideoList: DemoVideos,
		NextTime:  time.Now().Unix(),
	})
}
