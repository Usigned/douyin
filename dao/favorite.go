package dao

import (
	"log"
	"sync"
	"time"
)

type Favorite struct {
	Id int64
	//UserId   int64
	UserToken string
	VideoId   int64
	CreateAt  time.Time
}

type FavoriteDao struct {
}

var favoriteDao *FavoriteDao
var favoriteOnce sync.Once

// NewFavoriteDaoInstance Singleton
func NewFavoriteDaoInstance() *FavoriteDao {
	favoriteOnce.Do(
		func() {
			favoriteDao = &FavoriteDao{}
		})
	return favoriteDao
}

func (d *FavoriteDao) QueryVideoIdByToken(token string) ([]int64, error) {
	var ids []int64
	err := db.Where("user_token = ?", token).Find(&ids).Error
	if err != nil {
		return nil, err
	}
	return ids, nil
}

func (d *FavoriteDao) Save(favorite *Favorite) error {
	result := db.Create(&favorite)
	err := result.Error
	if err != nil {
		return err
	}
	return nil
}

func (d *FavoriteDao) Delete(favorite *Favorite) error {
	err := db.Where("user_token = ? AND video_id = ?", favorite.UserToken, favorite.VideoId).Delete(&Favorite{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (d *FavoriteDao) Total() (int64, error) {
	// 获取全部记录
	var count int64
	result := db.Table("comments").Count(&count)
	err := result.Error
	if err != nil {
		log.Fatal("total user err:" + err.Error())
		return -1, err
	}
	return count, nil
}
