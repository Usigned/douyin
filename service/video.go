package service

import (
	"github.com/Usigned/douyin/dao"
	"github.com/Usigned/douyin/entity"
	"sync"
	"time"
)

type VideoService struct {
}

var videoService *VideoService
var serviceOnce sync.Once

func NewVideoServiceInstance() *VideoService {
	serviceOnce.Do(
		func() {
			videoService = &VideoService{}
		})
	return videoService
}

func (s VideoService) FindVideoById(id int64) (*entity.Video, error) {
	return dao.NewVideoDaoInstance().QueryVideoById(id)
}

func (s VideoService) FindVideoBeforeTime(latestTime int64, limit int) ([]*entity.Video, error) {
	var t time.Time
	if latestTime == 0 {
		t = time.Now()
	} else {
		t = time.Unix(latestTime, 0)
	}
	return dao.NewVideoDaoInstance().QueryVideoBeforeTime(t, limit)
}
