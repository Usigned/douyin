package dao

import (
	"github.com/Usigned/douyin/entity"
	"gorm.io/gorm"
	"log"
	"sync"
	"time"
)

type Video struct {
	Id       int64
	AuthorId int64
	PlayUrl  string
	CoverUrl string
	Title    string
	CreateAt time.Time
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

func (*VideoDao) QueryVideoBeforeTime(time time.Time, limit int) ([]*Video, error) {
	var videos []*entity.Video
	err := db.Where("create_at < ?", time).Order("create_at DESC").Limit(limit).Find(&videos).Error

	if err != nil {
		log.Fatal("batch find video before time err:" + err.Error())
		return nil, err
	}
	return nil, nil
}
