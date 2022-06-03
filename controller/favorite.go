package controller

import (
	"douyin/dao"
	"douyin/entity"
	"douyin/pack"
	"douyin/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

var favoriteService = service.NewFavoriteServiceInstance()

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	token := c.Query("token")
	actionType := c.Query("action_type")
	if _, exist := usersLoginInfo[token]; !exist {
		if user, _ := favoriteService.FindUserByToken(token); user == nil {
			c.JSON(http.StatusOK, entity.Response{StatusCode: 1, StatusMsg: "User doesn't exist, Please Register!"})
			return
		}
	}
	var favorite = dao.Favorite{
		Id:        0,
		UserToken: token,
		VideoId:   videoId,
		CreateAt:  time.Now(),
	}
	if actionType == "1" {
		err := favoriteService.FavoriteAction(&favorite)
		if err != nil {
			c.JSON(http.StatusOK, entity.Response{StatusCode: 1, StatusMsg: "Favorite Failed, Please Retry!"})
		} else {
			c.JSON(http.StatusOK, entity.Response{StatusCode: 0, StatusMsg: "Thanks for your Favorite!"})
		}
	} else if actionType == "2" {
		err := favoriteService.FavoriteCancel(&favorite)
		if err != nil {
			c.JSON(http.StatusOK, entity.Response{StatusCode: 1, StatusMsg: "Cancel Failed, Please Retry!"})
		} else {
			c.JSON(http.StatusOK, entity.Response{StatusCode: 0, StatusMsg: "Please Favorite Next Time!"})
		}
	}
}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	token := c.Query("token")
	var videos []*entity.Video

	videos, err := favoriteService.FindVideoByToken(token)
	if err != nil {
		c.JSON(http.StatusOK, entity.Response{StatusCode: 1, StatusMsg: "Pull Favorite Failed!"})
	}

	c.JSON(http.StatusOK, VideoListResponse{
		Response: entity.Response{
			StatusCode: 0,
		},
		VideoList: pack.VideoPtrs(videos),
	})
}
