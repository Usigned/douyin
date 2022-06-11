package service

// TODO
import (
	"douyin/dao/mysql"
	"douyin/entity"
	"douyin/pack"
	"douyin/utils"
	"errors"
	"fmt"
	"strings"
	"sync"
)

type UserService struct {
}

var userService *UserService
var userOnce sync.Once
var usersLoginInfo = mysql.CopyULI()

func NewUserServiceInstance() *UserService {
	userOnce.Do(
		func() {
			userService = &UserService{}
		})
	return userService
}

func (s *UserService) UserInfo(id int64) (*entity.User, error) {
	// 查询用户信息
	userModel, err := mysql.NewUserDaoInstance().QueryUserById(id)
	if err != nil {
		return nil, err
	}

	// 包装用户信息
	user := pack.User(userModel)
	user.IsFollow = true
	return user, nil
}

func (s *UserService) FindUserByName(name string) (*entity.User, error) {
	user, err := mysql.NewUserDaoInstance().QueryUserByName(name)
	if err != nil || user == nil {
		return nil, err
	}
	return pack.User(user), nil
}

// AddUser 创建用户和token
func (s *UserService) AddUser(token string) error {
	fmt.Println("token:", token)
	parseToken, err := utils.ParseToken(token)

	if err != nil {
		return utils.Error{Msg: "invalid token"}
	}
	//maxId, _ := dao.NewUserDaoInstance().MaxId()
	//fmt.Println("maxId:", maxId)
	//userIdSequence := atomic.AddInt64(&maxId, 1)
	fmt.Println("666")
	newUser := &mysql.User{
		//Id: userIdSequence,
		Name:     parseToken.Username,
		Password: parseToken.Password,
	}
	fmt.Println("newUser:", newUser)
	err = mysql.NewUserDaoInstance().CreateUser(newUser)
	if err != nil {
		return err
	}

	// 创建token
	loginStatus := &mysql.LoginStatus{
		UserId: newUser.Id,
		Token:  token,
	}
	usersLoginInfo[loginStatus.Token] = *pack.User(newUser)
	err = mysql.NewLoginStatusDaoInstance().CreateLoginStatus(loginStatus)
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
	//password = utils.Md5(password)
	token, _ := utils.GenToken(username, utils.Md5(password))
	// 先查缓存 ..
	if _, exist := usersLoginInfo[token]; !exist {
		if user, _ := userService.FindUserByName(username); user == nil {
			err = s.AddUser(token)
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
	user, _ := s.FindUserByName(username)
	if user == nil {
		return nil, nil, utils.Error{Msg: "User doesn't exist, Please Register! "}
	}

	// 生成JWT Token
	//password = utils.Md5(password)
	token, _ := utils.GenToken(username, utils.Md5(password))

	usersLoginInfo[token] = *user
	// 密码校验
	result, _ := mysql.NewUserDaoInstance().QueryUserByToken(token)
	if result == nil {
		return nil, nil, utils.Error{Msg: "Password Wrong!"}
	}
	// 创建token
	loginStatus := &mysql.LoginStatus{
		UserId: user.Id,
		Token:  token,
		//Token:  utils.GenerateUUID(),
	}
	err := mysql.NewLoginStatusDaoInstance().CreateLoginStatus(loginStatus)
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
