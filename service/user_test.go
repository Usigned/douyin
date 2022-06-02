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
	userMap, err := userService.MFindUserById([]int64{1, 2})
	assert.Equal(t, err, nil)
	assert.NotEqual(t, userMap, nil)

	assert.Equal(t, len(userMap), 2)
	assert.Equal(t, userMap[1].Id, int64(1))
	assert.Equal(t, userMap[1].Name, "qing")
	assert.Equal(t, userMap[1].FollowerCount, int64(0))
	assert.Equal(t, userMap[1].FollowCount, int64(0))

	assert.Equal(t, userMap[2].Id, int64(2))
	assert.Equal(t, userMap[2].Name, "TestUser #1")
	assert.Equal(t, userMap[2].FollowerCount, int64(0))
	assert.Equal(t, userMap[2].FollowCount, int64(0))
}

func TestUserService_MFindUserById_Empty(t *testing.T) {
	users, err := userService.MFindUserById([]int64{-1})
	assert.Equal(t, err, nil)
	assert.NotEqual(t, users, nil)
	assert.Equal(t, len(users), 0)
}

func TestUserService_Register(t *testing.T) {
	err := userService.AddUser("lzq", "cdma1330")
	assert.Equal(t, err, nil)

	user, err := userService.FindUserByName("lzq")
	assert.Equal(t, err, nil)
	assert.NotEqual(t, user, nil)
}

func TestUserService_Register_Fail(t *testing.T) {
	err := userService.AddUser("lzq", "cdm1330")
	assert.NotEqual(t, err, nil)
}

func TestUserService_Login(t *testing.T) {
	id, token, err := userService.Login("TestUser #2", "123456")
	assert.Equal(t, err, nil)
	assert.NotEqual(t, id, nil)
	assert.NotEqual(t, token, nil)

	assert.Equal(t, *id, int64(7))
}

func TestUserService_Login_Fail_1(t *testing.T) {
	id, token, err := userService.Login("TestUser #2", "123456afdfaf")
	assert.NotEqual(t, err, nil)
	assert.Equal(t, id, nil)
	assert.Equal(t, token, nil)
}

func TestUserService_Login_Fail_2(t *testing.T) {
	id, token, err := userService.Login("TestUser #3", "123456")
	assert.NotEqual(t, err, nil)
	assert.Equal(t, id, nil)
	assert.Equal(t, token, nil)
}
