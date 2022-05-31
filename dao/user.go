package dao

import (
	"gorm.io/gorm"
	"log"
	"sync"
)

type User struct {
	Id            int64
	Name          string
	Password      string
	FollowCount   int64
	FollowerCount int64
	VideoCount    int64
	LikeCount     int64
}

type UserDao struct {
}

var userDao *UserDao
var userOnce sync.Once

// NewUserDaoInstance Singleton
func NewUserDaoInstance() *UserDao {
	userOnce.Do(
		func() {
			userDao = &UserDao{}
		})
	return userDao
}

func (*UserDao) QueryUserById(id int64) (*User, error) {
	var user User
	err := db.Where("id = ?", id).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		log.Fatal("find user by id err:" + err.Error())
		return nil, err
	}
	return &user, nil
}

func (*UserDao) MQueryUserById(ids []int64) ([]*User, error) {
	return nil, nil
}
