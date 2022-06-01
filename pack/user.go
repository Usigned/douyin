package pack

import (
	"github.com/Usigned/douyin/dao"
	"github.com/Usigned/douyin/entity"
)

// User if param is nil then return nil
func User(userModel *dao.User) *entity.User {
	if userModel != nil {
		return &entity.User{
			Id:            userModel.Id,
			Name:          userModel.Name,
			FollowCount:   userModel.FollowerCount,
			FollowerCount: userModel.FollowerCount,
		}
	}
	return nil
}

// MUser if param is nil then return empty map
func MUser(userModels map[int64]dao.User) map[int64]entity.User {
	if userModels != nil {
		var users = make(map[int64]entity.User, len(userModels))
		for id, userModel := range userModels {
			users[id] = *User(&userModel)
		}
		return users
	}
	return nil
}
