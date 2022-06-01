package service

import (
	"github.com/Usigned/douyin/dao"
	"github.com/Usigned/douyin/entity"
	"github.com/Usigned/douyin/pack"
	"sync"
)

type UserService struct {
}

var userService *UserService
var userOnce sync.Once

func NewUserServiceInstance() *UserService {
	userOnce.Do(
		func() {
			userService = &UserService{}
		})
	return userService
}

// FindUserById return nil if no user is found
func (s *UserService) FindUserById(id int64) (*entity.User, error) {
	userModel, err := dao.NewUserDaoInstance().QueryUserById(id)
	if err != nil {
		return nil, err
	}
	return pack.User(userModel), nil
}

// MFindUserById return empty map if no user is found
func (s *UserService) MFindUserById(ids []int64) (map[int64]entity.User, error) {
	userModels, err := dao.NewUserDaoInstance().MQueryUserById(ids)
	if err != nil {
		return nil, err
	}
	return pack.MUser(userModels), nil
}

func (s *UserService) FindTokenByUserId(id int64) (*string, error) {
	status, err := dao.NewLoginStatusDaoInstance().QueryByUserId(id)
	if err != nil || status == nil {
		return nil, err
	}
	return &status.Token, nil
}

func (s UserService) FindUserByToken(token string) (*entity.User, error) {
	// TODO
	return nil, nil
}
