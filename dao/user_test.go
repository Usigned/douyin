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

	userMap, err := userDao.MQueryUserById([]int64{1, 2})
	if err != nil {
		return
	}
	assert.Equal(t, err, nil)

	assert.Equal(t, userMap[1].Id, int64(1))
	assert.Equal(t, userMap[1].Name, "qing")
}

func TestUserDao_MQueryUserById_Empty(t *testing.T) {

	users, err := userDao.MQueryUserById([]int64{-1})
	assert.Equal(t, err, nil)
	assert.NotEqual(t, users, nil)
	assert.Equal(t, len(users), 0)
}

func TestUserDao_QueryUserByName(t *testing.T) {
	user, err := userDao.QueryUserByName("qing")
	assert.Equal(t, err, nil)
	assert.NotEqual(t, user, nil)
	assert.Equal(t, user.Id, int64(1))
}

func TestUserDao_IncreaseVideoCountByOne_NoSuchUser(t *testing.T) {
	err := userDao.IncreaseVideoCountByOne(-1)
	assert.NotEqual(t, err, nil)
}

func TestUserDao_IncreaseVideoCountByOne(t *testing.T) {
	err := userDao.IncreaseVideoCountByOne(1)
	assert.Equal(t, err, nil)
}
