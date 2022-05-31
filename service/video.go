package service

import (
	"github.com/Usigned/douyin/dao"
	"github.com/Usigned/douyin/entity"
	"github.com/Usigned/douyin/utils"
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
	var videoModel, err = dao.NewVideoDaoInstance().QueryVideoById(id)
	return utils.PackVideo(videoModel), err
}

func (s VideoService) FindVideoBeforeTime(latestTime int64, limit int) ([]*entity.Video, error) {
	var t time.Time
	if latestTime == 0 {
		t = time.Now()
	} else {
		t = time.UnixMilli(latestTime)
	}

	var videoModels, err = dao.NewVideoDaoInstance().QueryVideoBeforeTime(t, limit)

	if err != nil {
		return nil, err
	}

	return utils.MPackVideo(videoModels), nil
}
