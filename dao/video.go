package dao

import (
	"gorm.io/gorm"
	"log"
	"sync"
	"time"
)

type Video struct {
	Id            int64
	AuthorId      int64
	PlayUrl       string
	CoverUrl      string
	Title         string
	CreateAt      time.Time
	FavoriteCount int64
	CommentCount  int64
}

type VideoDao struct {
}

var videoDao *VideoDao
var videoOnce sync.Once

// NewVideoDaoInstance Singleton
func NewVideoDaoInstance() *VideoDao {
	videoOnce.Do(
		func() {
			videoDao = &VideoDao{}
		})
	return videoDao
}

// QueryVideoById will return nil if no user is found
func (*VideoDao) QueryVideoById(id int64) (*Video, error) {
	var video Video
	err := db.Where("id = ?", id).First(&video).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		log.Fatal("find video by id err:" + err.Error())
		return nil, err
	}
	return &video, nil
}

// QueryVideoBeforeTime will return empty array if no user is found
func (*VideoDao) QueryVideoBeforeTime(time time.Time, limit int) ([]*Video, error) {
	var videos []*Video
	err := db.Where("create_at < ?", time).Order("create_at DESC").Limit(limit).Find(&videos).Error

	if err != nil {
		log.Fatal("batch find video before time err:" + err.Error())
		return nil, err
	}
	return videos, nil
}

func (*VideoDao) CreateVideo(video *Video) error {
	return db.Create(&video).Error
}

func (*VideoDao) QueryVideoByAuthorId(authorId int64) ([]*Video, error) {
	var videos []*Video
	err := db.Where("author_id = ?", authorId).Find(&videos).Error
	if err != nil {
		log.Fatal("batch find video by author_id err:" + err.Error())
		return nil, err
	}
	return videos, nil
}

func (*VideoDao) UpdateCommentByID(id int64, count int64) error {
	err := db.Model(&Video{}).Where("id = ?", id).UpdateColumn("comment_count", count).Error
	if err != nil {
		return err
	}
	return nil
}
