package service

import (
	"douyin/dao"
)

var id int64

// FindFollowList 关注查询
func FindFollowList(query dao.Query) ([]dao.UserList, error) {
	id = query.UserId
	return dao.FindFollowList(id)
}

// FindFollowerList 粉丝查询
func FindFollowerList(query dao.Query) ([]dao.UserList, error) {
	id = query.UserId
	return dao.FindFollowerList(id)
}

// RelationAction 关注与取关操作
func RelationAction(change dao.Change) error {
	actionType := change.ActionType
	userId := id
	toUserId := change.ToUserId
	if actionType == 1 {
		//	关注
		return dao.FollowAction(userId, toUserId)
	} else if actionType == 2 {
		//	取关
		return dao.FollowerAction(userId, toUserId)
	}
	return nil
}
