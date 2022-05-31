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
	assert.Equal(t, videos[0].Id > videos[1].Id, true)
}

func TestVideoService_FindVideoAfterTime_Empty(t *testing.T) {
	users, err := videoService.FindVideoAfterTime(
		time.Date(1990, time.Month(12), 0, 0, 0, 0, 0, time.UTC).Unix(),
		utils.DefaultLimit)
	assert.Equal(t, err, nil)
	assert.NotEqual(t, users, nil)
	assert.Equal(t, len(users), 0)
}
