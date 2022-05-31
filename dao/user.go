package dao

// TODO

import (
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
	return nil, nil
}

func (*UserDao) MQueryUserById(ids []int64) ([]*User, error) {
	return nil, nil
}
