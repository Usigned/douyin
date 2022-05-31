package dao

import (
	"github.com/Usigned/douyin/utils"
	"github.com/go-playground/assert/v2"
	"testing"
	"time"
)

var _ = Init()

func TestVideoDao_QueryVideoById(t *testing.T) {
	video, err := videoDao.QueryVideoById(1)
	if err != nil {
		println(err.Error())
	}

	assert.Equal(t, video.Id, int64(1))
}

func TestVideoDao_MQueryVideoBeforeTime(t *testing.T) {
	videos, err := videoDao.MQueryVideoBeforeTime(
		time.Date(2022, time.Month(5), 30, 1, 5, 0, 0, time.UTC),
		utils.DefaultLimit)
	if err != nil {
		return
	}
	assert.Equal(t, len(videos), 1)
	assert.Equal(t, videos[0].Id, int64(2))
}

func TestVideoDao_CreateVideo(t *testing.T) {
	var video = &Video{
		AuthorId: 2,
		PlayUrl:  "12345678.com",
		CoverUrl: "cover1.com",
		Title:    "Test Video #1",
		CreateAt: time.Now(),
	}
	err := videoDao.CreateVideo(video)
	assert.Equal(t, err, nil)
}
