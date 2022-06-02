package service

// TODO
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
	// 查询用户信息
	userModel, err := dao.NewUserDaoInstance().QueryUserById(id)
	if err != nil {
		return nil, err
	}
	// 包装用户信息
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
	user, err := dao.NewUserDaoInstance().QueryUserByToken(token)
	if err != nil {
		return nil, err
	}
	return pack.User(user), err
}

func (s *UserService) FindUserByName(name string) (*entity.User, error) {
	// 查询用户信息
	userModel, err := dao.NewUserDaoInstance().QueryUserByName(name)
	if err != nil {
		return nil, err
	}

	// 包装用户信息
	return pack.User(userModel), nil
}

func (s *UserService) SaveUser(user *dao.User) (*entity.User, error) {
	err := dao.NewUserDaoInstance().Save(user)
	if err != nil {
		return nil, err
	}
	return pack.User(user), nil
}

func (s *UserService) TotalUser() (int64, error) {
	count, err := dao.NewUserDaoInstance().Total()
	if err != nil {
		return -1, err
	}
	return count, nil
}

func (s *UserService) LastId() (int64, error) {
	count, err := dao.NewUserDaoInstance().MaxId()
	if err != nil {
		return -1, err
	}
	return count, nil
}
