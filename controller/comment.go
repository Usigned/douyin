package controller

import (
	"douyin/entity"
	"douyin/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var commentService = service.NewCommentServiceInstance()

type CommentActionResponse struct {
	entity.Response
	Comment entity.Comment `json:"comment,omitempty"`
}

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	c.JSON(http.StatusOK, CommentActionFunc(
		c.Query("video_id"),
		c.Query("token"),
		c.Query("action_type"),
		c.Query("comment_id"),
		c.Query("comment_text"),
	))
}

func CommentActionFunc(videoId, token, actionType, commentId, text string) CommentActionResponse {
	vid, err := strconv.ParseInt(videoId, 10, 64)
	if err != nil {
		return ErrorCommentResponse(err)
	}
	if actionType == "1" {
		comment, err := commentService.Add(vid, token, text)
		if err != nil {
			return ErrorCommentResponse(err)
		}
		if comment == nil {
			return FailCommentResponse("Comments are not allowed to be empty! ")
		}
		return CommentActionResponse{
			Response: entity.Response{
				StatusCode: 0,
				StatusMsg:  "Add comment success! ",
			},
			Comment: *comment,
		}
	} else if actionType == "2" {
		cid, err := strconv.ParseInt(commentId, 10, 64)
		if err != nil {
			return ErrorCommentResponse(err)
		}
		comment, err := commentService.Withdraw(cid)
		if err != nil {
			return ErrorCommentResponse(err)
		}
		if comment == nil {
			return FailCommentResponse("Withdraw failed, Please try again later! ")
		}
		return CommentActionResponse{
			Response: entity.Response{
				StatusCode: 0,
				StatusMsg:  "Withdraw comment success! ",
			},
			Comment: *comment,
		}
	} else {
		return CommentActionResponse{
			Response: entity.Response{
				StatusCode: 1,
				StatusMsg:  "Service Wrong!",
			},
		}
	}
}

// ErrorCommentResponse 评论操作错误
func ErrorCommentResponse(err error) CommentActionResponse {
	return CommentActionResponse{
		Response: entity.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		},
	}
}

// FailCommentResponse 评论操作失败
func FailCommentResponse(msg string) CommentActionResponse {
	return CommentActionResponse{
		Response: entity.Response{
			StatusCode: -1,
			StatusMsg:  msg,
		},
	}
}
