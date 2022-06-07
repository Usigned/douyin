package service

import (
	"douyin/dao"
	"douyin/entity"
	"douyin/pack"
	"fmt"
	"sync"
)

type RelationService struct {
}

var relationService *RelationService
var relationOnce sync.Once

func NewRelationServiceInstance() *RelationService {
	relationOnce.Do(
		func() {
			relationService = &RelationService{}
		})
	return relationService
}

// Follow 关注操作
func (s *RelationService) Follow(userId, toUserId int64, token string) error {
	fmt.Println("当前用户：", userId)
	fmt.Println("登录用户：", toUserId)
	err := dao.NewRelationDaoInstance().FollowAction(userId, toUserId)
	if err != nil {
		return err
	}
	return nil
}

// WithdrawFollow 取关操作
func (s *RelationService) WithdrawFollow(userId, toUserId int64, token string) error {
	fmt.Println("当前用户：", userId)
	fmt.Println("登录用户：", toUserId)
	err := dao.NewRelationDaoInstance().WithdrawFollowAction(userId, toUserId)
	if err != nil {
		return err
	}
	return nil
}

// FollowList 关注查询
func (s *RelationService) FollowList(userId int64, token string) ([]*entity.User, error) {
	// 当前用户id : userId
	// 登录用户id : toUserId
	toUserId, err := dao.NewLoginStatusDaoInstance().QueryUserIdByToken(token)
	userListModels, err := dao.NewRelationDaoInstance().QueryFollowList(userId)
	if err != nil {
		return nil, err
	}
	users := pack.Users(userListModels)
	for _, user := range users {
		user.IsFollow = dao.NewRelationDaoInstance().IsFollow(toUserId, user.Id)
	}
	return users, nil
}

// FollowerList 粉丝查询
func (s *RelationService) FollowerList(userId int64, token string) ([]*entity.User, error) {
	// 当前用户id : userId
	// 登录用户id : toUserId
	toUserId, err := dao.NewLoginStatusDaoInstance().QueryUserIdByToken(token)
	userListModels, err := dao.NewRelationDaoInstance().QueryFollowerList(userId)
	if err != nil {
		return nil, err
	}
	users := pack.Users(userListModels)
	for _, user := range users {
		user.IsFollow = dao.NewRelationDaoInstance().IsFollow(toUserId, user.Id)
	}
	return users, nil
}
