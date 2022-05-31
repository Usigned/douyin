package service

import (
	"github.com/Usigned/douyin/dao"
	"github.com/go-playground/assert/v2"
	"testing"
)

var _ = dao.Init()
var _ = NewUserServiceInstance()

func TestUserService_FindUserById(t *testing.T) {
	user, err := userService.FindUserById(1)
	assert.Equal(t, err, nil)
	assert.NotEqual(t, user, nil)
	assert.Equal(t, user.Id, int64(1))
	assert.Equal(t, user.IsFollow, false)
}

func TestUserService_FindUserById_Nil(t *testing.T) {
	user, err := userService.FindUserById(-11)
	assert.Equal(t, err, nil)
	assert.Equal(t, user, nil)
}

func TestUserService_MFindUserById_NotEmpty(t *testing.T) {
	users, err := userService.MFindUserById([]int64{1, 2})
	assert.Equal(t, err, nil)
	assert.NotEqual(t, users, nil)

	//assert.Equal(t, len(users), 2)
	assert.Equal(t, users[0].Id, int64(1))
	assert.Equal(t, users[0].Name, "qing")
	assert.Equal(t, users[0].FollowerCount, int64(0))
	assert.Equal(t, users[0].FollowCount, int64(0))
}

func TestUserService_MFindUserById_Empty(t *testing.T) {
	users, err := userService.MFindUserById([]int64{-1})
	assert.Equal(t, err, nil)
	assert.NotEqual(t, users, nil)
	assert.Equal(t, len(users), 0)
}
