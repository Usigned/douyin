package service

import (
	"github.com/Usigned/douyin/dao"
	"github.com/Usigned/douyin/entity"
	"github.com/Usigned/douyin/pack"
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

func (s *VideoService) FindVideoById(id int64) (*entity.Video, error) {
	videoModel, err := dao.NewVideoDaoInstance().QueryVideoById(id)
	if err != nil {
		return nil, err
	}

	if videoModel == nil {
		return nil, nil
	}

	userModel, err := dao.NewUserDaoInstance().QueryUserById(videoModel.AuthorId)
	if err != nil {
		return nil, err
	}

	user := pack.User(userModel)
	video := pack.Video(videoModel)

	video.Author = *user
	return video, nil
}

// FindVideoAfterTime return video info packed with user info
func (s *VideoService) FindVideoAfterTime(latestTime int64, limit int) ([]*entity.Video, error) {
	var t time.Time
	if latestTime == 0 {
		t = time.Now()
	} else {
		t = time.UnixMilli(latestTime)
	}

	videoModels, err := dao.NewVideoDaoInstance().QueryVideoBeforeTime(t, limit)
	if err != nil {
		return nil, err
	}

	userModels, err := dao.NewUserDaoInstance().MQueryUserById(pack.MAuthorId(videoModels))
	if err != nil {
		return nil, err
	}

	users := pack.MUser(userModels)
	videos := pack.MVideo(videoModels)

	for i, video := range videos {
		video.Author = *users[i]
	}

	return videos, nil
}

func (s *VideoService) FindVideoByAuthorId(authorId int64) ([]*entity.Video, error) {
	// invalid authorId
	if authorId <= 0 {
		return nil, nil
	}

	videoModels, err := dao.NewVideoDaoInstance().QueryVideoByAuthorId(authorId)
	if err != nil {
		return nil, err
	}
	userModels, err := dao.NewUserDaoInstance().MQueryUserById(pack.MAuthorId(videoModels))
	if err != nil {
		return nil, err
	}

	users := pack.MUser(userModels)
	videos := pack.MVideo(videoModels)

	for i, video := range videos {
		video.Author = *users[i]
	}

	return videos, nil
}
