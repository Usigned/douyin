package dao

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

var _ = Init()

//func TestUserDao_CreateUser(t *testing.T) {
//	var user = &User{
//		Name:     "TestUser #1",
//		Password: "123456",
//	}
//	err := userDao.CreateUser(user)
//	assert.Equal(t, err, nil)
//}

func TestUserDao_QueryUserById_NotNil(t *testing.T) {

	user, err := userDao.QueryUserById(1)

	assert.Equal(t, err, nil)

	assert.Equal(t, user.Id, int64(1))
	assert.Equal(t, user.Name, "qing")
}

func TestUserDao_QueryUserById_Nil(t *testing.T) {

	user, err := userDao.QueryUserById(-1)
	assert.Equal(t, err, nil)
	assert.Equal(t, user, nil)
}

func TestUserDao_MQueryUserById_NotEmpty(t *testing.T) {

	users, err := userDao.MQueryUserById([]int64{1, 2})
	if err != nil {
		return
	}
	assert.Equal(t, err, nil)

	assert.Equal(t, users[0].Id, int64(1))
	assert.Equal(t, users[0].Name, "qing")
}

func TestUserDao_MQueryUserById_Empty(t *testing.T) {

	users, err := userDao.MQueryUserById([]int64{-1})
	assert.Equal(t, err, nil)
	assert.Equal(t, users, []*User{})
}
