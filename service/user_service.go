package service

// TODO
import (
	"douyin/dao"
	"douyin/entity"
	"douyin/pack"
	"douyin/utils"
	"errors"
	"strings"
	"sync"
)

type UserService struct {
}

var userService *UserService
var userOnce sync.Once
var usersLoginInfo = dao.CopyULI()

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
	user := pack.User(userModel)
	user.IsFollow = true
	return user, nil
}

func (s *UserService) FindUserByName(name string) (*entity.User, error) {
	user, err := dao.NewUserDaoInstance().QueryUserByName(name)
	if err != nil || user == nil {
		return nil, err
	}
	return pack.User(user), nil
}

func (s *UserService) FindUserByToken(token string) (*entity.User, error) {
	user, err := dao.NewUserDaoInstance().QueryUserByToken(token)
	if err != nil || user == nil {
		return nil, err
	}
	return pack.User(user), nil
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
	status, err := dao.NewLoginStatusDaoInstance().QueryTokenByUserId(user.Id)
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

// AddUser 创建用户和token
func (s *UserService) AddUser(username, password string) error {
	// 用户注册
	password = utils.Md5(password)
	token := "<" + username + "><" + password + ">"
	//userIdSequence, _ := dao.NewUserDaoInstance().MaxId()
	//atomic.AddInt64(&userIdSequence, 1)
	newUser := &dao.User{
		//Id:       userIdSequence,
		Name:     username,
		Password: password,
	}
	err := dao.NewUserDaoInstance().CreateUser(newUser)
	if err != nil {
		return err
	}

	// 创建token
	loginStatus := &dao.LoginStatus{
		UserId: newUser.Id,
		Token:  token,
		//Token:  utils.GenerateUUID(),
	}
	usersLoginInfo[loginStatus.Token] = *pack.User(newUser)
	err = dao.NewLoginStatusDaoInstance().CreateLoginStatus(loginStatus)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) Register(username, password string) error {
	// 用户输入验证
	err := InfoVerify(username, password)
	if err != nil {
		return err
	}
	token := "<" + username + "><" + password + ">"
	// 先查缓存 ..
	if _, exist := usersLoginInfo[token]; !exist {
		if user, _ := userService.FindUserByName(username); user == nil {
			err = s.AddUser(username, password)
			if err != nil {
				return utils.Error{Msg: "User register failed, Please retry for a minute!"}
			}
			return err
		}
	}
	return utils.Error{Msg: "User already exist, don't register again!"}
}

func (s *UserService) Login(username, password string) (*int64, *string, error) {
	// 用户校验
	password = utils.Md5(password)
	token := "<" + username + "><" + password + ">"

	user, _ := s.FindUserByName(username)
	if user == nil {
		return nil, nil, utils.Error{Msg: "User doesn't exist, Please Register! "}
	}
	usersLoginInfo[token] = *user
	// 密码校验
	result, _ := dao.NewUserDaoInstance().QueryUserByToken(token)
	if result == nil {
		return nil, nil, utils.Error{Msg: "Password Wrong!"}
	}
	// 创建token
	loginStatus := &dao.LoginStatus{
		UserId: user.Id,
		Token:  token,
		//Token:  utils.GenerateUUID(),
	}
	err := dao.NewLoginStatusDaoInstance().CreateLoginStatus(loginStatus)
	if err != nil {
		return nil, nil, err
	}
	return &user.Id, &token, nil
}

func InfoVerify(username string, password string) error {
	if Check(username) {
		return errors.New("Please Check Username!\nThe length is controlled within 4-32 characters, and <, >, \\is not allowed")
	}
	if Check(password) {
		return errors.New("Please Check Password!\nThe length is controlled within 4-32 characters, and <, >, \\is not allowed")
	}
	return nil
}

func Check(str string) bool {
	length := len(str)
	if length < 4 || length > 32 {
		return true
	}
	if strings.Contains(str, "<") || strings.Contains(str, ">") ||
		strings.Contains(str, "/") || strings.Contains(str, "\\") {
		return true
	}
	return false
}
