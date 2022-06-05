package service

// TODO
import (
	"douyin/dao"
	"douyin/entity"
	"douyin/pack"
	"douyin/utils"
	"sync"
	"sync/atomic"
	"time"
)

type FavoriteService struct {
}

var favoriteService *FavoriteService
var favoriteOnce sync.Once

func NewFavoriteServiceInstance() *FavoriteService {
	favoriteOnce.Do(
		func() {
			favoriteService = &FavoriteService{}
		})
	return favoriteService
}

func (s *FavoriteService) FindUserByToken(token string) (*entity.User, error) {
	user, err := dao.NewUserDaoInstance().QueryUserByToken(token)
	if err != nil {
		return nil, err
	}
	return pack.User(user), err
}

func (s *FavoriteService) FindVideosByToken(token string) ([]*entity.Video, error) {
	// invalid token
	if token == "" {
		return nil, nil
	}
	videoIds, err := dao.NewFavoriteDaoInstance().QueryVideoIdByToken(token)
	if err != nil {
		return nil, err
	}
	var videos []*entity.Video
	for _, id := range videoIds {
		video, _ := NewVideoServiceInstance().FindVideoById(id)
		//video.IsFavorite = true
		videos = append(videos, video)
	}
	return videos, nil
}

func (s *FavoriteService) TotalComment() (int64, error) {
	count, err := dao.NewFavoriteDaoInstance().Total()
	if err != nil {
		return -1, err
	}
	return count, nil
}

func (s *FavoriteService) LastId() (int64, error) {
	count, err := dao.NewFavoriteDaoInstance().MaxId()
	if err != nil {
		return count, err
	}
	return count, nil
}

func (s *FavoriteService) Add(videoId int64, token string) error {
	// 先查缓存 ..
	if _, exist := usersLoginInfo[token]; !exist {
		user, _ := dao.NewUserDaoInstance().QueryUserByToken(token)
		if user == nil {
			return utils.Error{Msg: "User doesn't exist, Please Register! "}
		}
		usersLoginInfo[token] = *pack.User(user)
	}
	// 点赞
	favoriteIdSequence, _ := favoriteService.LastId()
	atomic.AddInt64(&favoriteIdSequence, 1)
	newFavorite := &dao.Favorite{
		Id:        favoriteIdSequence,
		UserToken: token,
		VideoId:   videoId,
		CreateAt:  time.Now(),
	}
	err := dao.NewFavoriteDaoInstance().Save(newFavorite)
	if err != nil {
		return err
	}
	return nil
}

func (s *FavoriteService) Withdraw(videoId int64, token string) error {
	// 删除评论
	err := dao.NewFavoriteDaoInstance().Delete(videoId, token)
	if err != nil {
		return err
	}
	return nil
}

func (s *FavoriteService) FavoriteList(token string) ([]*entity.Video, error) {
	return s.FindVideosByToken(token)
}
