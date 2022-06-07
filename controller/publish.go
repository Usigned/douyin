package controller

import (
	"douyin/entity"
	"douyin/service"
	"douyin/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

// 是否自动生成封面，需要配置环境，默认为否
var useGeneratedCover = utils.UseGeneratedCover

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
	fileName := utils.GenerateUUID()
	videoFileName := fmt.Sprintf("%s%s", fileName, ext)
	coverName := fmt.Sprintf("%s%s", fileName, ".jpg")

	// 判断文件夹是否存在
	var dir = "./publish/"
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		os.Mkdir(dir, os.ModePerm)
	}

	saveFile := filepath.Join(dir, videoFileName)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		return ErrorResponse(err)
	}
	//生成视频信息
	// TODO 目前是数据库硬编码ip:port，后续可改成动态
	playUrl := utils.VideoUrlPrefix + videoFileName
	//封面
	var coverUrl string
	if useGeneratedCover {
		coverUrl = utils.VideoUrlPrefix + coverName
		// generate video cover
		coverFilePath := filepath.Join(dir, coverName)
		utils.ReadFrameAsJpeg(saveFile, 1, coverFilePath)
	} else {
		coverUrl = utils.DefaultCoverUrl
	}

	err = service.NewVideoServiceInstance().Publish(token, playUrl, coverUrl, title)
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
