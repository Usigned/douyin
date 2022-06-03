package service

import (
	"github.com/Usigned/douyin/dao"
	"github.com/Usigned/douyin/utils"
	"github.com/go-playground/assert/v2"
	"testing"
	"time"
)

var _ = dao.Init()
var _ = NewVideoServiceInstance()

func TestVideoService_FindVideoById(t *testing.T) {
	video, err := videoService.FindVideoById(1)
	assert.Equal(t, err, nil)
	assert.Equal(t, video.Author.Id, int64(1))
	assert.Equal(t, video.Author.Name, "qing")
}

func TestVideoService_FindVideoById_Nil(t *testing.T) {
	video, err := videoService.FindVideoById(-11)
	assert.Equal(t, err, nil)
	assert.Equal(t, video, nil)
}

func TestVideoService_FindVideoAfterTime(t *testing.T) {
	videos, err := videoService.FindVideoAfterTime(0, utils.DefaultLimit)
	assert.Equal(t, err, nil)
	assert.NotEqual(t, videos, nil)

	//assert.Equal(t, len(users), 2)
	assert.Equal(t, videos[0].Id, int64(2))
	assert.Equal(t, videos[0].Author.Id, int64(2))
}

func TestVideoService_MFindVideoAfterTime_Empty(t *testing.T) {
	users, err := videoService.FindVideoAfterTime(
		time.Date(1990, time.Month(12), 0, 0, 0, 0, 0, time.UTC).Unix(),
		utils.DefaultLimit)
	assert.Equal(t, err, nil)
	assert.NotEqual(t, users, nil)
	assert.Equal(t, len(users), 0)
}

func TestVideoService_MFindVideoByAuthorId(t *testing.T) {
	videos, err := videoService.FindVideoByAuthorId(1)
	assert.Equal(t, err, nil)
	assert.NotEqual(t, videos, nil)
	assert.Equal(t, len(videos), 1)
}

func TestVideoService_Publish(t *testing.T) {
	user, err := dao.NewUserDaoInstance().QueryUserById(12)
	videoCount := user.VideoCount
	err = videoService.Publish("2b6383b5-7078-4800-8678-8e82256a85fc", "test_video_url", "test_cover_url", "Test Video #123")
	assert.Equal(t, err, nil)
	user, err = dao.NewUserDaoInstance().QueryUserById(12)
	assert.Equal(t, err, nil)
	assert.NotEqual(t, user, nil)
	assert.Equal(t, user.VideoCount, videoCount+1)
}
