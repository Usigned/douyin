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
	//var latestTime int64
	////var token string
	//
	//t, _ := c.GetQuery("latest_time")
	//latestTime, _ = strconv.ParseInt(t, 10, 64)
	////token, _ = c.GetQuery("token")
	//
	//videos, err := videoService.FindVideoBeforeTime(latestTime, utils.DefaultLimit)
	//if err != nil {
	//	c.JSON(http.StatusOK, FeedResponse{
	//		Response: entity.Response{
	//			StatusCode: 1,
	//			StatusMsg:  err.Error(),
	//		},
	//		VideoList: nil,
	//		NextTime:  time.Now().Unix(),
	//	})
	//	return
	//}
	//
	//var videoList = make([]entity.Video, 0, len(videos))
	//for _, video := range videos {
	//	videoVo := entity.Video{}
	//	err := videoVo.fromEntity(video)
	//	if err != nil {
	//		c.JSON(http.StatusOK, FeedResponse{
	//			Response: entity.Response{
	//				StatusCode: 1,
	//				StatusMsg:  err.Error(),
	//			},
	//			VideoList: nil,
	//			NextTime:  time.Now().Unix(),
	//		})
	//		return
	//	}
	//	videoList = append(videoList, videoVo)
	//}

	c.JSON(http.StatusOK, FeedResponse{
		Response: entity.Response{
			StatusCode: 0,
			StatusMsg:  "success",
		},
		VideoList: DemoVideos,
		NextTime:  time.Now().Unix(),
	})
}
