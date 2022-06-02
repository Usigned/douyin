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
	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, entity.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, PublishFunc(token, data, c))

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

func PublishFunc(token string, data *multipart.FileHeader, c *gin.Context) entity.Response {
	// TODO
	// 1. check token
	// 2. store video file
	// 3. add video

	err := service.NewVideoServiceInstance().Publish(token, "afdafa", "13131", "title")
	if err != nil {
		return ErrorResponse(err)
	}
	return ErrorResponse(utils.Error{Msg: "upload fail"})
}

func ErrorResponse(err error) entity.Response {
	return entity.Response{
		StatusCode: 1,
		StatusMsg:  err.Error(),
	}
}
