package controller

import (
	"douyin/entity"
	"douyin/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var favoriteService = service.NewFavoriteServiceInstance()

type FavoriteActionResponse struct {
	entity.Response
}

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	c.JSON(http.StatusOK, FavoriteActionFunc(
		c.Query("video_id"),
		c.Query("token"),
		c.Query("action_type"),
	))
}

func FavoriteActionFunc(videoId, token, actionType string) FavoriteActionResponse {
	vid, err := strconv.ParseInt(videoId, 10, 64)
	if err != nil {
		return ErrorFavoriteResponse(err)
	}

	if actionType == "1" {
		err = favoriteService.Add(vid, token)
		if err != nil {
			return ErrorFavoriteResponse(err)
		}
		return FavoriteActionResponse{
			Response: entity.Response{
				StatusCode: 0,
				StatusMsg:  "Thanks for your favorite! ",
			},
		}
	} else if actionType == "2" {
		err := favoriteService.Withdraw(vid, token)
		if err != nil {
			return ErrorFavoriteResponse(err)
		}
		return FavoriteActionResponse{
			Response: entity.Response{
				StatusCode: 0,
				StatusMsg:  "Please Favorite Next Time! ",
			},
		}
	} else {
		return FavoriteActionResponse{
			Response: entity.Response{
				StatusCode: 1,
				StatusMsg:  "Service Wrong!",
			},
		}
	}
}

func ErrorFavoriteResponse(err error) FavoriteActionResponse {
	return FavoriteActionResponse{
		Response: entity.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		},
	}
}
