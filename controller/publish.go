package controller

import (
	"fmt"
	"github.com/Usigned/douyin/entity"
	"github.com/Usigned/douyin/service"
	"github.com/Usigned/douyin/utils"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"net/http"
	"path/filepath"
)

// Publish check token then save upload file to public directory TODO
func Publish(c *gin.Context) {
	token := c.PostForm("token")
	title := c.PostForm("title")
	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, entity.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, PublishFunc(token, title, data, c))
}

// PublishFunc TODO
func PublishFunc(token, title string, data *multipart.FileHeader, c *gin.Context) entity.Response {
	//检查文件是否为空
	if data == nil {
		return ErrorResponse(utils.Error{Msg: "empty data file"})
	}
	//检查后缀名
	ext := filepath.Ext(data.Filename)
	if ext != ".mp4" {
		return ErrorResponse(utils.Error{Msg: "unsupported file extension"})
	}
	//存文件
	filepath.Base(data.Filename)
	filename := fmt.Sprintf("%s%s", utils.GenerateUUID(), ext)
	saveFile := filepath.Join("./public/", filename)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		return ErrorResponse(err)
	}
	//生成视频信息
	// TODO 目前是数据库硬编码域名，后续可改成动态
	// http://127.0.0.1/douyin/video/filename
	playUrl := filepath.Join(utils.VideoUrlPrefix, filename)
	coverUrl := utils.DefaultCoverUrl

	err := service.NewVideoServiceInstance().Publish(token, playUrl, coverUrl, title)
	if err != nil {
		return ErrorResponse(err)
	}
	return entity.Response{
		StatusCode: 0,
		StatusMsg:  "success",
	}
}

func ErrorResponse(err error) entity.Response {
	return entity.Response{
		StatusCode: 1,
		StatusMsg:  err.Error(),
	}
}
