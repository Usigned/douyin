package controller

import (
	"douyin/entity"
	"douyin/pack"
	"douyin/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CommentListResponse struct {
	entity.Response
	CommentList []entity.Comment `json:"comment_list,omitempty"`
}

// CommentList ..
func CommentList(c *gin.Context) {
	c.JSON(http.StatusOK, CommentListFunc(
		c.Query("video_id"),
		c.Query("token"),
	))
}

func CommentListFunc(videoID, token string) CommentListResponse {
	// TODO 使用token进行鉴权
	if videoID == "" {
		return ErrorCommentListResponse(utils.Error{Msg: "empty token or user_id"})
	}
	vid, err := strconv.ParseInt(videoID, 10, 64)
	if err != nil {
		return ErrorCommentListResponse(err)
	}
	comments, err := commentService.LoadComments(vid)
	if err != nil {
		ErrorCommentListResponse(err)
	}
	return CommentListResponse{
		Response: entity.Response{
			StatusCode: 0,
			StatusMsg:  "Load comments success!",
		},
		CommentList: pack.CommentsPtrs(comments),
	}
}

func ErrorCommentListResponse(err error) CommentListResponse {
	return CommentListResponse{
		Response: entity.Response{
			StatusCode: -1,
			StatusMsg:  err.Error(),
		},
		CommentList: nil,
	}
}
