package service

import (
	"douyin/dao"
	"douyin/entity"
	"douyin/utils"
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
	//id, err := dao.NewLoginStatusDaoInstance().QueryUserIdByToken(t)
	err := dao.FollowAction(userId, toUserId)
	if err != nil {
		return err
	}
	return nil
}

// Follower 取关操作
func (s *RelationService) Follower(userId, toUserId int64, token string) error {
	err := dao.FollowerAction(userId, toUserId)
	if err != nil {
		return err
	}
	return nil
}

// FollowList 关注查询
func (s *RelationService) FollowList(userId int64, token string) ([]entity.User, error) {
	// 其他用户id : userId
	// 当前用户id : toUserId
	toUserId, err := dao.NewLoginStatusDaoInstance().QueryUserIdByToken(token)
	if err != nil {
		return nil, err
	}
	if toUserId == -1 {
		return nil, utils.Error{Msg: "user not exist"}
	}
	return dao.FindFollowList(toUserId, userId)
}

// FollowerList 粉丝查询
func (s *RelationService) FollowerList(userId int64, token string) ([]entity.User, error) {
	toUserId, err := dao.NewLoginStatusDaoInstance().QueryUserIdByToken(token)
	if err != nil {
		return nil, err
	}
	if toUserId == -1 {
		return nil, utils.Error{Msg: "user not exist"}
	}
	return dao.FindFollowerList(toUserId, userId)
}
