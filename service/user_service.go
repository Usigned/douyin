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
	"sync/atomic"
)

type UserService struct {
}

var usersLoginInfo = map[string]entity.User{
	"zhangleidouyin": {
		Id:            1,
		Name:          "zhanglei",
		FollowCount:   10,
		FollowerCount: 5,
		IsFollow:      true,
	},
	"qingcdma1330": {
		Id:            1,
		Name:          "Qing",
		FollowCount:   100,
		FollowerCount: 5000,
		IsFollow:      false,
	},
}

var userService *UserService
var userOnce sync.Once

func CopyULI() map[string]entity.User {
	return usersLoginInfo
}

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

func (s *UserService) FindUserByName(name string) (*entity.User, error) {
	user, err := dao.NewUserDaoInstance().QueryUserByName(name)
	if err != nil || user == nil {
		return nil, err
	}
	return pack.User(user), nil
}

// AddUser 创建用户和token
func (s *UserService) AddUser(username, password string) error {
	// 用户注册
	password = utils.Md5(password)
	userIdSequence, _ := dao.NewUserDaoInstance().MaxId()
	atomic.AddInt64(&userIdSequence, 1)
	newUser := &dao.User{
		Id:       userIdSequence,
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
		Token:  utils.GenerateUUID(),
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
	//// 校验用户名及密码是否合法
	//err := InfoVerify(username, password)
	//if err != nil {
	//	return nil, nil, err
	//}
	// 用户校验
	password = utils.Md5(password)
	token := "<" + username + "><" + password + ">"
	// 先查询缓存 ..
	user, _ := dao.NewUserDaoInstance().QueryUserByName(username)
	if _, exist := usersLoginInfo[token]; !exist {
		if user == nil {
			return nil, nil, utils.Error{Msg: "User doesn't exist, Please Register! "}
		}
		usersLoginInfo[token] = *pack.User(user)
	}
	// 密码校验
	result, _ := dao.NewUserDaoInstance().QueryUserByToken(token)
	if result == nil {
		return nil, nil, utils.Error{Msg: "Password Wrong!"}
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
