package utils

import (
	"github.com/Usigned/douyin/dao"
	"github.com/Usigned/douyin/entity"
)

func PackUser(userModel *dao.User) *entity.User {
	return nil
}

func MPackUser(userModels []*dao.User) []*entity.User {
	return nil
}

func MPackAuthorId(videoModels []*dao.Video) []int64 {
	return []int64{1, 1, 1, 1}
}

func PackVideo(videoModel *dao.Video) *entity.Video {
	return nil
}

func MPackVideo(videoModels []*dao.Video) []*entity.Video {
	return nil
}
