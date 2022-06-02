package service

// TODO
import (
	"github.com/Usigned/douyin/dao"
	"github.com/Usigned/douyin/entity"
	"github.com/Usigned/douyin/pack"
	"sync"
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

func (s *FavoriteService) FindVideoByToken(token string) ([]*entity.Video, error) {
	videoIds, err := dao.NewFavoriteDaoInstance().QueryVideoIdByToken(token)
	if err != nil {
		return nil, err
	}
	var videos []*entity.Video
	for _, id := range videoIds {
		video, _ := NewVideoServiceInstance().FindVideoById(id)
		videos = append(videos, video)
	}
	return videos, nil
}

func (s *FavoriteService) FavoriteAction(favorite *dao.Favorite) error {
	err := dao.NewFavoriteDaoInstance().Save(favorite)
	if err != nil {
		return err
	}
	return nil
}

func (s *FavoriteService) FavoriteCancel(favorite *dao.Favorite) error {
	err := dao.NewFavoriteDaoInstance().Delete(favorite)
	if err != nil {
		return err
	}
	return nil
}

func (s *FavoriteService) TotalComment() (int64, error) {
	count, err := dao.NewFavoriteDaoInstance().Total()
	if err != nil {
		return -1, err
	}
	return count, nil
}
