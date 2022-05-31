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

// MUser if param is nil then return empty array
func MUser(userModels []*dao.User) []*entity.User {
	if userModels != nil {
		var users = make([]*entity.User, 0, len(userModels))
		for _, userModel := range userModels {
			users = append(users, User(userModel))
		}
		return users
	}
	return nil
}
