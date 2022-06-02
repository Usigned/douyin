package service

import (
	"github.com/Usigned/douyin/dao"
	"github.com/Usigned/douyin/entity"
	"github.com/Usigned/douyin/pack"
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

// FindTokenByUserName 根据用户名查找token，如果用户不存在则抛出异常，如果用户存在,token不存在则创建新token
func (s *UserService) FindTokenByUserName(name string) (*string, error) {
	// 查询用户是否存在
	user, err := dao.NewUserDaoInstance().QueryUserByName(name)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, utils.Error{Msg: "user not exist"}
	}
	// 查询现有的token
	status, err := dao.NewLoginStatusDaoInstance().QueryByUserId(user.Id)
	if err != nil {
		return nil, err
	}

	// token不存在
	if status == nil {
		status := &dao.LoginStatus{
			UserId: user.Id,
			Token:  utils.GenerateUUID(),
		}
		err := dao.NewLoginStatusDaoInstance().CreateLoginStatus(status)
		if err != nil {
			return nil, err
		}
	}
	return &status.Token, nil
}

func (s *UserService) FindUserByName(name string) (*entity.User, error) {
	user, err := dao.NewUserDaoInstance().QueryUserByName(name)
	if err != nil || user == nil {
		return nil, err
	}
	return pack.User(user), nil
}

// AddUser 创建用户和token
func (s *UserService) AddUser(username, password string) error {
	if username == "" {
		return utils.Error{Msg: "Invalid username"}
	}

	// 创建用户
	user, err := dao.NewUserDaoInstance().QueryUserByName(username)
	if err != nil {
		return err
	}
	if user != nil {
		return utils.Error{Msg: "username already been used"}
	}

	password = utils.Md5(password)
	user = &dao.User{
		Name:     username,
		Password: password,
	}
	err = dao.NewUserDaoInstance().CreateUser(user)
	if err != nil {
		return err
	}

	// 创建token
	loginStatus := &dao.LoginStatus{
		UserId: user.Id,
		Token:  utils.GenerateUUID(),
	}
	err = dao.NewLoginStatusDaoInstance().CreateLoginStatus(loginStatus)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) Register(username, password string) error {
	return s.AddUser(username, password)
}

func (s *UserService) Login(username, password string) (*int64, *string, error) {
	// 查询用户是否存在
	user, err := dao.NewUserDaoInstance().QueryUserByName(username)
	if err != nil {
		return nil, nil, err
	}
	if user == nil {
		return nil, nil, utils.Error{Msg: "user not exist"}
	}

	// 校验密码
	if utils.Md5(password) != user.Password {
		return nil, nil, utils.Error{Msg: "wrong password"}
	}

	// 查询现有的token
	status, err := dao.NewLoginStatusDaoInstance().QueryByUserId(user.Id)
	if err != nil {
		return nil, nil, err
	}

	// token不存在
	if status == nil {
		status := &dao.LoginStatus{
			UserId: user.Id,
			Token:  utils.GenerateUUID(),
		}
		err := dao.NewLoginStatusDaoInstance().CreateLoginStatus(status)
		if err != nil {
			return nil, nil, err
		}
	}
	return &user.Id, &status.Token, nil
}
