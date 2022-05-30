package controller

import (
	"github.com/Usigned/douyin/entity"
	"github.com/Usigned/douyin/service"
	"github.com/Usigned/douyin/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

var (
	userService  = service.NewUserServiceInstance()
	videoService = service.NewVideoServiceInstance()
)

type FeedResponse struct {
	Response
	VideoList []Video `json:"video_list"`
	NextTime  int64   `json:"next_time"`
}

// Feed use userService and videoService to query data
func Feed(c *gin.Context) {
	var latestTime int64
	//var token string

	t, _ := c.GetQuery("latest_time")
	latestTime, _ = strconv.ParseInt(t, 10, 64)
	//token, _ = c.GetQuery("token")

	videos, err := videoService.FindVideoBeforeTime(latestTime, utils.DefaultLimit)
	if err != nil {
		c.JSON(http.StatusOK, FeedResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
			VideoList: nil,
			NextTime:  time.Now().Unix(),
		})
		return
	}

	var videoList = make([]Video, 0, len(videos))
	for _, video := range videos {
		videoVo := Video{}
		err := videoVo.fromEntity(video)
		if err != nil {
			c.JSON(http.StatusOK, FeedResponse{
				Response: Response{
					StatusCode: 1,
					StatusMsg:  err.Error(),
				},
				VideoList: nil,
				NextTime:  time.Now().Unix(),
			})
			return
		}
		videoList = append(videoList, videoVo)
	}

	c.JSON(http.StatusOK, FeedResponse{
		Response: Response{
			StatusCode: 0,
			StatusMsg:  "success",
		},
		VideoList: videoList,
		NextTime:  time.Now().Unix(),
	})
}

func (videoVo *Video) fromEntity(video *entity.Video) error {
	videoVo.Id = video.Id
	videoVo.CoverUrl = video.CoverUrl
	videoVo.PlayUrl = video.PlayUrl
	videoVo.Title = video.Title

	// add author info
	user, err := userService.FindUserById(video.AuthorId)
	if err != nil {
		return err
	}
	var author = User{}
	err = author.fromEntity(user)
	if err != nil {
		return err
	}
	videoVo.Author = author

	// TODO
	videoVo.FavoriteCount = 100
	videoVo.IsFavorite = true
	videoVo.CommentCount = 100

	return nil
}

func (userVo *User) fromEntity(user *entity.User) error {
	userVo.Id = user.Id
	userVo.Name = user.Name

	// TODO
	userVo.FollowCount = 100
	userVo.FollowerCount = 100
	userVo.IsFollow = false

	return nil
}
