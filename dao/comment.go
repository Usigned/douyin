package dao

// TODO

import (
	"log"
	"sync"
)

type Comment struct {
	Id       int64
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
	userOnce.Do(
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

func (d *CommentDao) QueryCommentById(id int64) (*Comment, error) {
	return nil, nil
}

func (d *CommentDao) QueryCommentByName(name string) (*Comment, error) {
	return nil, nil
}

func (d *CommentDao) Save(comment *Comment) error {
	result := db.Create(&comment)
	err := result.Error
	if err != nil {
		return err
	}
	return nil
}

func (d *CommentDao) DeleteCommentById(id int64) error {
	result := db.Where("id = ?", id).Delete(&Comment{})
	err := result.Error
	if err != nil {
		return err
	}
	return nil
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
