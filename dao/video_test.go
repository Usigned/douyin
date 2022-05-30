package dao

import (
	"fmt"
	"github.com/go-playground/assert/v2"
	"testing"
	"time"
)

func TestQuery(t *testing.T) {
	err := Init()
	if err != nil {
		println(err.Error())
	}

	videoDao := NewVideoDaoInstance()
	video, err := videoDao.QueryVideoById(4)
	if err != nil {
		println(err.Error())
	}

	assert.Equal(t, video.Id, int64(4))

	videos, err := videoDao.QueryVideoBeforeTime(time.Now(), 10)
	if err != nil {
		println(err.Error())
	}
	assert.Equal(t, len(videos), 2)
	for _, video := range videos {
		fmt.Printf("%#v\n", video)
	}
}

func TestParseTime(t *testing.T) {
	fmt.Printf("%#v\n", time.Unix(time.Now().Unix(), 0))
}
