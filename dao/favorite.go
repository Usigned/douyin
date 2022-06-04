package dao

import (
	"log"
	"sync"
	"time"
)

type Favorite struct {
	Id        int64
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
	err := db.Select("video_id").Table("favorites").Where("user_token = ?", token).Find(&ids).Error
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

func (d *FavoriteDao) Delete(videoId int64, token string) error {
	err := db.Where("user_token = ? AND video_id = ?", token, videoId).Delete(&Favorite{}).Error
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

func (d *FavoriteDao) MaxId() (int64, error) {
	// 获取全部记录
	var lastRec *Comment
	result := db.Table("favorites").Last(&lastRec)
	err := result.Error
	if err != nil {
		//log.Fatal("max id err:" + err.Error())
		return 0, err
	}
	return lastRec.Id, nil
}
