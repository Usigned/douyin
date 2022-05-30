package dao

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestUserDao_QueryUserById(t *testing.T) {
	err := Init()
	if err != nil {
		println(err.Error())
	}

	userDao := NewUserDaoInstance()
	user, err := userDao.QueryUserById(1)
	if err != nil {
		return
	}

	assert.Equal(t, user.Id, int64(1))
	assert.Equal(t, user.Name, "TestUser")
}

func TestUserDao_MQueryUserById(t *testing.T) {
	err := Init()
	if err != nil {
		println(err.Error())
	}

	userDao := NewUserDaoInstance()
	users, err := userDao.MQueryUserById([]int64{1, 2})
	if err != nil {
		return
	}

	assert.Equal(t, users[1].Id, int64(1))
	assert.Equal(t, users[1].Name, "TestUser")

	assert.Equal(t, users[2].Id, int64(2))
	assert.Equal(t, users[2].Name, "TestUser")
}
