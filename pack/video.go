package pack

import (
	"douyin/dao"
	"douyin/entity"
)

func AuthorIds(videoModels []*dao.Video) []int64 {
	if videoModels != nil {
		var ids = make([]int64, 0, len(videoModels))
		for _, videoModel := range videoModels {
			ids = append(ids, videoModel.AuthorId)
		}
		return ids
	}
	return []int64{}
}

func Video(videoModel *dao.Video) *entity.Video {
	if videoModel != nil {
		return &entity.Video{
			Id:            videoModel.Id,
			Author:        entity.User{},
			PlayUrl:       videoModel.PlayUrl,
			CoverUrl:      videoModel.CoverUrl,
			Title:         videoModel.Title,
			FavoriteCount: videoModel.FavoriteCount,
			CommentCount:  videoModel.FavoriteCount,
		}
	}
	return nil
}

func Videos(videoModels []*dao.Video) []*entity.Video {
	if videoModels != nil {
		var videos = make([]*entity.Video, 0, len(videoModels))
		for _, model := range videoModels {
			videos = append(videos, Video(model))
		}
		return videos
	}
	return nil
}

func VideoPtrs(videoPtrs []*entity.Video) []entity.Video {
	if videoPtrs != nil {
		var videos = make([]entity.Video, len(videoPtrs))
		for i, ptr := range videoPtrs {
			videos[i] = *ptr
		}
		return videos
	}
	return []entity.Video{}
}
