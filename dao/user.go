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

// QueryUserById will return nil if no user is found
func (*UserDao) QueryUserById(id int64) (*User, error) {
	var user *User
	err := db.Where("id = ?", id).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		log.Fatal("find user by id err:" + err.Error())
		return nil, err
	}
	return user, nil
}

// MQueryUserById will return empty array if no user is found
func (*UserDao) MQueryUserById(ids []int64) ([]*User, error) {
	var users []*User
	err := db.Where("id in (?)", ids).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (*UserDao) CreateUser(user *User) error {
	return db.Create(&user).Error
}
