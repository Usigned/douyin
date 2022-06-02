package dao

import (
	"gorm.io/gorm"
	"log"
	"sync"
)

type LoginStatus struct {
	Id     int64
	UserId int64
	Token  string
}

type LoginStatusDao struct {
}

var loginStatusDao *LoginStatusDao
var loginStatusOnce sync.Once

// NewLoginStatusDaoInstance Singleton
func NewLoginStatusDaoInstance() *LoginStatusDao {
	loginStatusOnce.Do(
		func() {
			loginStatusDao = &LoginStatusDao{}
		})
	return loginStatusDao
}

func (*LoginStatusDao) QueryByUserId(userId int64) (*LoginStatus, error) {
	var loginStatus *LoginStatus
	err := db.Where("user_id = ?", userId).First(&loginStatus).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		log.Fatal("find loginStatus by user_id err:" + err.Error())
		return nil, err
	}
	return loginStatus, nil
}

func (*LoginStatusDao) CreateLoginStatus(loginStatus *LoginStatus) error {
	return db.Create(&loginStatus).Error
}
