package service

import (
	"douyin/dao"
	"douyin/entity"
	"douyin/pack"
	"douyin/utils"
	"strconv"
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

func (s *VideoService) Feed(latestTime int64, token string, limit int) (*int64, []*entity.Video, error) {
	var t time.Time
	if latestTime == 0 {
		t = time.Now()
	} else {
		t = time.UnixMilli(latestTime)
	}

	videoModels, err := dao.NewVideoDaoInstance().QueryVideoBeforeTime(t, limit)
	if err != nil {
		return nil, nil, err
	}

	authorIds := pack.AuthorIds(videoModels)

	userModelMap, err := dao.NewUserDaoInstance().MQueryUserById(authorIds)
	if err != nil {
		return nil, nil, err
	}

	userMap := pack.MUser(userModelMap)

	// 获取当前用户
	curUserId, err := dao.NewLoginStatusDaoInstance().QueryUserIdByToken(token)
	if err != nil {
		return nil, nil, err
	}

	if curUserId != -1 {
		for uid := range userMap {
			userMap[uid].IsFollow = dao.NewRelationDaoInstance().IsFollow(curUserId, uid)
		}
	}

	videos := pack.Videos(videoModels)

	for i, video := range videos {
		video.Author = *userMap[authorIds[i]]

		commentCount, _, err := dao.NewCommentDaoInstance().QueryCommentByVideoId(video.Id)
		if err != nil {
			return nil, nil, err
		}
		video.CommentCount = commentCount

		favoriteCount, err := dao.NewFavoriteDaoInstance().QueryFavoriteByVideoId(video.Id)
		if err != nil {
			return nil, nil, err
		}
		video.FavoriteCount = favoriteCount
		video.IsFavorite = dao.NewFavoriteDaoInstance().QueryFavoriteByUserToken(video.Id, token)
	}

	var nextTime int64
	if len(videoModels) > 0 {
		nextTime = videoModels[len(videoModels)-1].CreateAt.UnixMilli()
	} else {
		nextTime = time.Now().UnixMilli()
	}
	println("latest time: " + strconv.FormatInt(latestTime, 10))
	println("next time: " + strconv.FormatInt(nextTime, 10))

	return &nextTime, videos, nil
}

func (s *VideoService) PublishList(authorId int64) ([]*entity.Video, error) {
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
		video.Author = *userMap[authorIds[i]]
	}

	return videos, nil
}

func (s VideoService) Publish(token, playUrl, coverUrl, title string) error {
	if playUrl == "" || coverUrl == "" || title == "" {
		return utils.Error{Msg: "参数不能为空"}
	}
	// 查询用户
	userId, err := dao.NewLoginStatusDaoInstance().QueryUserIdByToken(token)
	if err != nil {
		return err
	}
	if userId == -1 {
		return utils.Error{Msg: "user not exist"}
	}

	// 保存 video
	videoModel := dao.Video{
		AuthorId:      userId,
		PlayUrl:       playUrl,
		CoverUrl:      coverUrl,
		Title:         title,
		CreateAt:      time.Now(),
		FavoriteCount: 0,
		CommentCount:  0,
	}
	err = dao.NewVideoDaoInstance().CreateVideo(&videoModel)
	if err != nil {
		return err
	}
	// 用户的视频数增加
	err = dao.NewUserDaoInstance().IncreaseVideoCountByOne(userId)
	if err != nil {
		return err
	}
	return nil
}
