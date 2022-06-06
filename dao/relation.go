package dao

import (
	"douyin/entity"
	"fmt"
	"gorm.io/gorm"
)

type AttentionDao struct {
}

type Attention struct {
	Id       int64
	UserId   int64 `json:"user_id,omitempty"`
	ToUserId int64 `json:"to_user_id,omitempty"`
	IsFollow bool  `json:"is_follow,omitempty"`
}

func (v Attention) TableName() string {
	return "attention"
}

// FindFollowList 查询关注者的id、姓名、关注数、粉丝数、与当前用户的关系
func FindFollowList(id, idB int64) ([]entity.User, error) {
	var followList []entity.User
	//查询该用户关注的人中被自己关注的人
	if err := db.Debug().Table("users").Select("users.id,users.name,users.follow_count,users.follower_count,"+
		"t.is_follow").Joins("inner join (select * from attention where user_id = ? and"+
		" to_user_id in (select to_user_id from attention where user_id = ?)) t on t.to_user_id=users.id", id, idB).Find(&followList).Error; err != nil {
		return nil, err
	}

	fmt.Println("关注列表", followList)
	var unfollowList []entity.User
	//查询该用户关注的人中未被自己关注的人
	if err := db.Debug().Table("users").Select("users.id,users.name,users.follow_count,"+
		"users.follower_count").Joins("inner join (select * from attention where user_id = ? and"+
		" to_user_id not in (select to_user_id from attention where user_id = ?)) t on t.to_user_id=users.id", id, idB).Find(&unfollowList).Error; err != nil {
		return nil, err
	}
	fmt.Println("未关注列表", unfollowList)
	followList = append(followList, unfollowList...)
	return followList, nil
}

// FindFollowerList 查询粉丝的id、姓名、关注数、粉丝数、与当前用户的关系
func FindFollowerList(id, idB int64) ([]entity.User, error) {
	var followerList []entity.User
	//查询该用户的粉丝中被自己关注的人
	if err := db.Debug().Table("users").Select("users.id,users.name,users.follow_count,users.follower_count,"+
		"t.is_follow").Joins("inner join (select * from attention where to_user_id = ? and"+
		" user_id in (select to_user_id from attention where user_id = ?)) t on t.user_id=users.id", id, idB).Find(&followerList).Error; err != nil {
		return nil, err
	}
	fmt.Println("未关注列表", followerList)
	//查询该用户的粉丝中未被自己关注的人
	var unfollowerList []entity.User
	if err := db.Debug().Table("users").Select("users.id,users.name,users.follow_count,"+
		"users.follower_count").Joins("inner join (select * from attention where to_user_id = ? and"+
		" user_id not in (select to_user_id from attention where user_id = ?)) t on t.user_id=users.id", id, idB).Find(&unfollowerList).Error; err != nil {
		return nil, err
	}
	fmt.Println("未关注列表", unfollowerList)
	//if err := db.Debug().Table("users").Select("users.id,users.name,users.follow_count,users.follower_count,attention.is_follow").Joins("inner join attention on attention.user_id=users.id").Where("attention.to_user_id=?", id).Find(&followerList).Error; err != nil {
	//	return nil, err
	//}
	followerList = append(followerList, unfollowerList...)
	return followerList, nil
}

// FollowAction 关注操作
func FollowAction(idA, idB int64) error {
	//对关系表的操作
	relation := Attention{
		UserId:   idA,
		ToUserId: idB,
		IsFollow: true,
	}
	db.Debug().Create(&relation)
	//对用户表的操作
	var userA entity.User
	//Update(“follow_count”, gorm.Expr(“follow_count + ?”, 1))
	db.Debug().Where("id = ?", idA).First(&userA).Update("follow_count", gorm.Expr("follow_count + ?", 1))
	var userB entity.User
	db.Debug().Where("id = ?", idB).First(&userB).Update("follower_count", gorm.Expr("follower_count + ?", 1))

	return nil
}

// FollowerAction 取关操作
func FollowerAction(idA, idB int64) error {
	//对关系表的操作

	var relation Attention
	if err := db.Debug().Where("user_id=? and to_user_id =?", idA, idB).Delete(&relation).Error; err != nil {
		fmt.Println(relation)
		return err
	}
	//对用户表的操作
	var userA entity.User
	db.Debug().Where("id = ?", idA).First(&userA).Update("follow_count", gorm.Expr("follow_count - ?", 1))
	var userB entity.User
	db.Debug().Where("id = ?", idB).First(&userB).Update("follower_count", gorm.Expr("follower_count - ?", 1))
	return nil
}
