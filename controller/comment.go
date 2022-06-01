package controller

import (
	"fmt"
	"github.com/Usigned/douyin/dao"
	"github.com/Usigned/douyin/entity"
	"github.com/Usigned/douyin/pack"
	"github.com/Usigned/douyin/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"sync/atomic"
	"time"
)

var commentService = service.NewCommentServiceInstance()

type CommentListResponse struct {
	entity.Response
	CommentList []entity.Comment `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	entity.Response
	Comment entity.Comment `json:"comment,omitempty"`
}

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	token := c.Query("token")
	actionType := c.Query("action_type")
	// 先查缓存
	if user, exist := usersLoginInfo[token]; exist {
		if actionType == "1" {
			text := c.Query("comment_text")
			date := time.Now().Format("01-02")

			commentIdSequence, _ := commentService.LastId()
			fmt.Println("id is:", commentIdSequence)
			atomic.AddInt64(&commentIdSequence, 1)

			testComment := entity.Comment{
				Id:         commentIdSequence,
				User:       user,
				Content:    text,
				CreateDate: date,
			}
			var curComment = dao.Comment{
				Id:       commentIdSequence,
				UserName: user.Name,
				Content:  text,
				CreateAt: time.Now().Format("01-02"),
			}

			commentService.CommentAction(&curComment)

			c.JSON(http.StatusOK, CommentActionResponse{Response: entity.Response{StatusCode: 0},
				Comment: testComment})
			return
		} else if actionType == "2" {
			// 删除评论
			id, _ := strconv.ParseInt(c.Query("comment_id"), 10, 64)
			date := time.Now().Format("01-02")
			testComment := entity.Comment{
				Id:         id,
				User:       user,
				CreateDate: date,
			}
			fmt.Println("删除测试", testComment)
			commentService.CommentDelete(id)
			c.JSON(http.StatusOK, CommentActionResponse{Response: entity.Response{StatusCode: 0},
				Comment: testComment})
			return
		}
		c.JSON(http.StatusOK, entity.Response{StatusCode: 0})
	} else {
		// BUG：客户端用户处于登录状态，数据库删除了该用户信息，但是客户端的用户没有下线，功能异常
		c.JSON(http.StatusOK, entity.Response{StatusCode: 1, StatusMsg: "User doesn't exist, Please Register!"})
	}
}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	comments, err := commentService.LoadComments()
	if err != nil {
		c.JSON(http.StatusOK, entity.Response{StatusCode: 1, StatusMsg: "Loading comments failed!"})
	} else {
		c.JSON(http.StatusOK, CommentListResponse{
			Response:    entity.Response{StatusCode: 0},
			CommentList: pack.CommentsPtrs(comments),
		})
	}

}
