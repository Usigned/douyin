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

// Feed 新视频流接口
func (s *VideoService) Feed(latestTime int64, limit int) ([]*entity.Video, error) {
	return s.FindVideoAfterTime(latestTime, limit)
}

// FindVideoAfterTime return video info packed with user info
// 老接口，新接口使用Feed
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

	authorIds := pack.AuthorIds(videoModels)

	userModelMap, err := dao.NewUserDaoInstance().MQueryUserById(authorIds)
	if err != nil {
		return nil, err
	}

	userMap := pack.MUser(userModelMap)
	videos := pack.Videos(videoModels)

	for i, video := range videos {
		video.Author = userMap[authorIds[i]]
	}

	return videos, nil
}

// PublishList 新发布列表接口
func (s *VideoService) PublishList(authorId int64) ([]*entity.Video, error) {
	return s.FindVideoByAuthorId(authorId)
}

// FindVideoByAuthorId 老接口，新接口使用PublishList
func (s *VideoService) FindVideoByAuthorId(authorId int64) ([]*entity.Video, error) {
	// invalid authorId
	if authorId <= 0 {
		return nil, nil
	}

	videoModels, err := dao.NewVideoDaoInstance().QueryVideoByAuthorId(authorId)
	if err != nil {
		return nil, err
	}
	authorIds := pack.AuthorIds(videoModels)

	userModelMap, err := dao.NewUserDaoInstance().MQueryUserById(authorIds)
	if err != nil {
		return nil, err
	}

	userMap := pack.MUser(userModelMap)
	videos := pack.Videos(videoModels)

	for i, video := range videos {
		video.Author = userMap[authorIds[i]]
	}

	return videos, nil
}
