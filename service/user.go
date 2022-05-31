package service

import (
	"github.com/Usigned/douyin/dao"
	"github.com/Usigned/douyin/entity"
	"github.com/Usigned/douyin/utils"
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

func (s *UserService) FindUserById(id int64) (*entity.User, error) {
	userModel, err := dao.NewUserDaoInstance().QueryUserById(id)
	if err != nil {
		return nil, err
	}
	return utils.PackUser(userModel), nil
}

func (s *UserService) MFindUserById(ids []int64) ([]*entity.User, error) {
	userModels, err := dao.NewUserDaoInstance().MQueryUserById(ids)
	if err != nil {
		return nil, err
	}
	return utils.MPackUser(userModels), nil
}
