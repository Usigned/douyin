package controller

import (
	"douyin/entity"
	"douyin/pack"
	"douyin/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type FavoriteListResponse struct {
	entity.Response
	VideoList []entity.Video `json:"video_list"`
}

// FavoriteList ..
func FavoriteList(c *gin.Context) {
	c.JSON(http.StatusOK, FavoriteListFunc(
		c.Query("token"),
	))
}

func FavoriteListFunc(token string) FavoriteListResponse {
	// TODO 使用token进行鉴权
	if token == "" {
		return ErrorFavoriteListResponse(utils.Error{Msg: "empty token or user_id"})
	}
	videos, err := favoriteService.FavoriteList(token)
	if err != nil {
		ErrorFavoriteListResponse(err)
	}
	return FavoriteListResponse{
		Response: entity.Response{
			StatusCode: 0,
			StatusMsg:  "Load Favorites success!",
		},
		VideoList: pack.VideoPtrs(videos),
	}
}

func ErrorFavoriteListResponse(err error) FavoriteListResponse {
	return FavoriteListResponse{
		Response: entity.Response{
			StatusCode: -1,
			StatusMsg:  err.Error(),
		},
		VideoList: nil,
	}
}
