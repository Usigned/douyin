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

// MQueryVideoBeforeTime will return empty array if no user is found
func (*VideoDao) MQueryVideoBeforeTime(time time.Time, limit int) ([]*Video, error) {
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
