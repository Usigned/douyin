package dao

// TODO

import (
	"gorm.io/gorm"
	"log"
	"regexp"
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

func (*UserDao) CreateUser(user *User) error {
	return db.Create(&user).Error
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

// MQueryUserById will return empty array if no user is found
func (*UserDao) MQueryUserById(ids []int64) (map[int64]User, error) {
	var users []*User
	err := db.Where("id in (?)", ids).Find(&users).Error
	if err != nil {
		return nil, err
	}
	var userMap = make(map[int64]User, len(users))
	for _, user := range users {
		id := user.Id
		userMap[id] = *user
	}
	return userMap, nil
}

func (d *UserDao) MQueryUserByName(names []string) (map[string]User, error) {
	var users []*User
	err := db.Where("name in (?)", names).Find(&users).Error
	if err != nil {
		return nil, err
	}
	var userMap = make(map[string]User, len(users))
	for _, user := range users {
		userMap[user.Name] = *user
	}
	return userMap, nil
}

func (*UserDao) QueryUserByName(name string) (*User, error) {
	var user *User
	err := db.Where("name = ?", name).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (*UserDao) QueryUserByToken(token string) (*User, error) {
	var users *User //实例化对象
	re, err := regexp.Compile("[A-Za-z0-9_\\-\u4e00-\u9fa5]+")
	if err != nil {
		return nil, err
	}
	name := re.FindAllString(token, 2)[0]
	password := re.FindAllString(token, 2)[1]
	err = db.Debug().Where("name = ? and password = ?", name, password).First(&users).Error
	if err == gorm.ErrRecordNotFound {
		return nil, err
	}
	if err != nil {
		//fmt.Println("record not found!")
		return nil, err
	}

	return users, nil
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
		return 0, err
	}
	return lastRec.Id, nil
}

func (*UserDao) IncreaseVideoCountByOne(id int64) error {
	var user *User
	err := db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return err
	}
	user.VideoCount = user.VideoCount + 1
	return db.Save(&user).Error
}
