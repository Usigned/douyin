package service

import (
	"github.com/Usigned/douyin/dao"
	"github.com/Usigned/douyin/entity"
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
	return dao.NewUserDaoInstance().QueryUserById(id)
}

func (s *UserService) MFindUserById(ids []int64) (map[int64]*entity.User, error) {
	return dao.NewUserDaoInstance().MQueryUserById(ids)
}
