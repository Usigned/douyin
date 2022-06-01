package dao

// TODO

import (
	"gorm.io/gorm"
	"log"
	"sync"
)

type User struct {
	Id            int64
	Name          string `gorm:"unique;not null"`
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
	user := new(User) //实例化对象
	result := db.Where("id = ?", id).First(&user)
	err := result.Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		log.Fatal("find user by id err:" + err.Error())
		return nil, err
	}
	return user, nil
}

func (*UserDao) QueryUserByName(name string) (*User, error) {
	var user *User //实例化对象
	result := db.Where("name = ?", name).First(&user)
	err := result.Error
	if err == gorm.ErrRecordNotFound {
		return nil, err
	}
	if err != nil {
		//fmt.Println("record not found!")
		return nil, err
	}
	return user, nil
}

func (*UserDao) MQueryUserById(ids []int64) ([]*User, error) {
	return nil, nil
}

func (*UserDao) Save(user *User) error {
	result := db.Create(&user)
	err := result.Error
	if err != nil {
		return err
	}
	return nil
}

func (*UserDao) Total() (int64, error) {
	// 获取全部记录
	var count int64
	result := db.Table("users").Count(&count)
	err := result.Error
	if err != nil {
		log.Fatal("total user err:" + err.Error())
		return -1, err
	}
	return count, nil
}

func (*UserDao) MaxId() (int64, error) {
	// 获取全部记录
	var lastRec *User
	result := db.Table("users").Last(&lastRec)
	err := result.Error
	if err != nil {
		log.Fatal("max id err:" + err.Error())
		return 0, err
	}
	return lastRec.Id, nil
}
