package dao

import (
	"fmt"
	"gorm.io/gorm"
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

func (d *FavoriteDao) QueryFavoriteByVideoId(videoID int64) (int64, error) {
	var favoriteCount int64
	result := db.Table("videos").Select("favorite_count").Where("id = ?", videoID).Find(&favoriteCount)
	err := result.Error
	if err != nil {
		return 0, err
	}
	return favoriteCount, nil
}

func (d *FavoriteDao) QueryVideoIdByToken(token string) ([]int64, error) {
	var ids []int64
	err := db.Select("video_id").Table("favorites").Where("user_token = ?", token).Find(&ids).Error
	if err != nil {
		return nil, err
	}
	return ids, nil
}

func (d *FavoriteDao) QueryFavoriteByUserToken(videoId int64, token string) bool {
	err := db.Where("video_id = ? AND user_token = ?", videoId, token).First(&Favorite{}).Error
	if err != nil {
		return false
	}
	return true
}

func (d *FavoriteDao) Save(favorite *Favorite) error {
	result := db.Create(&favorite)
	err := result.Error
	if err != nil {
		return err
	}

	err = db.Debug().Model(&Video{}).Where("id = ?", favorite.VideoId).Update("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (d *FavoriteDao) Delete(videoId int64, token string) error {
	err := db.Where("user_token = ? AND video_id = ?", token, videoId).Delete(&Favorite{}).Error
	if err != nil {
		return err
	}

	err = db.Debug().Model(&Video{}).Where("id = ?", videoId).Update("favorite_count", gorm.Expr("favorite_count - ?", 1)).Error
	if err != nil {
		fmt.Println(err)
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
