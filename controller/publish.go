package controller

import (
	"fmt"
	"github.com/Usigned/douyin/entity"
	"github.com/Usigned/douyin/pack"
	"github.com/Usigned/douyin/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"strconv"
)

type VideoListResponse struct {
	entity.Response
	VideoList []entity.Video `json:"video_list"`
}

type InvalidParameterError struct {
	msg string
}

func (e InvalidParameterError) Error() string {
	return e.msg
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	// TODO
	token := c.PostForm("token")

	if _, exist := usersLoginInfo[token]; !exist {
		c.JSON(http.StatusOK, entity.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, entity.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	filename := filepath.Base(data.Filename)
	user := usersLoginInfo[token]
	finalName := fmt.Sprintf("%d_%s", user.Id, filename)
	saveFile := filepath.Join("./public/", finalName)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, entity.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, entity.Response{
		StatusCode: 0,
		StatusMsg:  finalName + " uploaded successfully",
	})
}

// PublishList /douyin/publish
func PublishList(c *gin.Context) {
	c.JSON(http.StatusOK, publishListFunc(c.Query("token"), c.Query("user_id")))
}

func publishListFunc(token, userId string) VideoListResponse {
	// TODO 使用token进行鉴权
	if token == "" || userId == "" {
		return ErrorVideoListResponse(InvalidParameterError{msg: "empty token or user_id"})
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
