package dao

import (
	"github.com/Usigned/douyin/entity"
	"gorm.io/gorm"
	"log"
	"sync"
)

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

func (*UserDao) QueryUserById(id int64) (*entity.User, error) {
	var user entity.User
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

func (*UserDao) MQueryUserById(ids []int64) (map[int64]*entity.User, error) {
	var users []*entity.User
	err := db.Where("id in (?)", ids).Find(&users).Error
	if err != nil {
		log.Fatal("batch find user by id err:" + err.Error())
		return nil, err
	}
	userMap := make(map[int64]*entity.User)
	for _, user := range users {
		userMap[user.Id] = user
	}
	return userMap, nil
}
