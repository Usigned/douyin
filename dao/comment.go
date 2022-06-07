package dao

// TODO

import (
	"log"
	"sync"
)

type Comment struct {
	Id       int64
	VideoId  int64
	UserName string
	Content  string
	CreateAt string
}

type CommentDao struct {
}

var commentDao *CommentDao
var commentOnce sync.Once

// NewCommentDaoInstance Singleton
func NewCommentDaoInstance() *CommentDao {
	commentOnce.Do(
		func() {
			commentDao = &CommentDao{}
		})
	return commentDao
}

func (d *CommentDao) QueryAllComment() ([]*Comment, error) {
	// 获取全部记录
	var comments []*Comment
	err := db.Find(&comments).Error
	if err != nil {
		//log.Fatal("batch find video by author_id err:" + err.Error())
		return nil, err
	}
	return comments, nil
}

func (d *CommentDao) QueryCommentById(id int64) ([]*Comment, error) {
	var comments []*Comment
	err := db.Where("id = ?", id).Find(&comments).Error
	if err != nil {
		log.Fatal("batch find video by author_id err:" + err.Error())
		return nil, err
	}
	return comments, nil
}

// QueryCommentByVideoId 添加返回结果int64
func (d *CommentDao) QueryCommentByVideoId(videoID int64) (int64, []*Comment, error) {
	var comments []*Comment
	result := db.Where("video_id = ?", videoID).Order("id DESC").Find(&comments)
	err := result.Error
	if err != nil {
		return 0, nil, err
	}
	return result.RowsAffected, comments, nil
}

func (d *CommentDao) QueryCommentByName(name string) (*Comment, error) {
	return nil, nil
}

func (d *CommentDao) Save(comment *Comment) (*Comment, error) {
	result := db.Create(&comment)
	err := result.Error
	if err != nil {
		return nil, err
	}
	return comment, nil
}

func (d *CommentDao) DeleteById(id int64) (*Comment, error) {
	var comment *Comment
	result := db.Where("id = ?", id).Delete(&comment)
	err := result.Error
	if err != nil {
		return nil, err
	}
	return comment, nil
}

func (d *CommentDao) Total() (int64, error) {
	// 获取全部记录
	var count int64
	result := db.Table("comments").Count(&count)
	err := result.Error
	if err != nil {
		log.Fatal("total user err:" + err.Error())
		return -1, err
	}
	return count, nil
}

func (d *CommentDao) TotalById(id int64) (int64, error) {
	// 获取全部记录
	var count int64
	result := db.Table("comments").Where("video_id = ?", id).Count(&count)
	err := result.Error
	if err != nil {
		log.Fatal("total user err:" + err.Error())
		return -1, err
	}
	return count, nil
}

func (d *CommentDao) MaxId() (int64, error) {
	// 获取全部记录
	var lastRec *Comment
	result := db.Table("comments").Last(&lastRec)
	err := result.Error
	if err != nil {
		//log.Fatal("max id err:" + err.Error())
		return 0, err
	}
	return lastRec.Id, nil
}
