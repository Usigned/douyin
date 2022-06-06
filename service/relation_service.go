package service

import (
	"douyin/dao"
	"douyin/entity"
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
func (s *RelationService) Follow(userId, toUserId int64) error {
	err := dao.FollowAction(userId, toUserId)
	if err != nil {
		return err
	}
	return nil
}

// Follower 取关操作
func (s *RelationService) Follower(userId, toUserId int64) error {
	err := dao.FollowerAction(userId, toUserId)
	if err != nil {
		return err
	}
	return nil
}

// FollowList 关注查询
func (s *RelationService) FollowList(userId int64) ([]entity.User, error) {
	list, err := dao.FindFollowList(userId)
	if err != nil {
		return nil, err
	}
	return list, nil
}

// FollowerList 粉丝查询
func (s *RelationService) FollowerList(userId int64) ([]entity.User, error) {
	list, err := dao.FindFollowerList(userId)
	if err != nil {
		return nil, err
	}
	return list, nil
}
