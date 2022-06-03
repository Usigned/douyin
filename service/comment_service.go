package service

// TODO
import (
	"douyin/dao"
	"douyin/entity"
	"douyin/pack"
	"sync"
)

type CommentService struct {
}

var commentService *CommentService
var commentOnce sync.Once

func NewCommentServiceInstance() *CommentService {
	commentOnce.Do(
		func() {
			commentService = &CommentService{}
		})
	return commentService
}

func (s *CommentService) LoadComments() ([]*entity.Comment, error) {
	// 查询用户信息
	commentModels, err := dao.NewCommentDaoInstance().QueryAllComment()
	if err != nil {
		return nil, err
	}

	// 包装用户信息
	return pack.Comments(commentModels), nil
}

func (s *CommentService) FindUserById(id int64) (*entity.Comment, error) {
	// 查询用户信息
	commentModel, err := dao.NewCommentDaoInstance().QueryCommentById(id)
	if err != nil {
		return nil, err
	}

	// 包装用户信息
	return pack.Comment(commentModel), nil
}

func (s *CommentService) FindCommentByName(name string) (*entity.Comment, error) {
	// 查询用户信息
	commentModel, err := dao.NewCommentDaoInstance().QueryCommentByName(name)
	if err != nil {
		return nil, err
	}

	// 包装用户信息
	return pack.Comment(commentModel), nil
}

func (s *CommentService) MFindCommentById(ids []int64) ([]*entity.Comment, error) {
	return nil, nil
}

func (s *CommentService) CommentAction(comment *dao.Comment) error {
	err := dao.NewCommentDaoInstance().Save(comment)
	if err != nil {
		return err
	}
	commentCount, err := s.TotalComment()
	dao.NewVideoDaoInstance().UpdateCommentByID(comment.VideoId, commentCount)
	return nil
}

func (s *CommentService) CommentDelete(id int64) error {
	videoId, err := dao.NewCommentDaoInstance().DeleteCommentById(id)
	if err != nil {
		return err
	}
	commentCount, err := s.TotalComment()
	dao.NewVideoDaoInstance().UpdateCommentByID(videoId, commentCount)
	return nil
}

func (s *CommentService) TotalComment() (int64, error) {
	count, err := dao.NewCommentDaoInstance().Total()
	if err != nil {
		return -1, err
	}
	return count, nil
}

func (s *CommentService) LastId() (int64, error) {
	count, err := dao.NewCommentDaoInstance().MaxId()
	if err != nil {
		return count, err
	}
	return count, nil
}
