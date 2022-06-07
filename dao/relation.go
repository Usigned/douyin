package dao

import (
	"douyin/entity"
	"douyin/utils"
	"fmt"
	"gorm.io/gorm"
	"sync"
)

type Attention struct {
	Id       int64
	UserId   int64 `json:"user_id,omitempty"`
	ToUserId int64 `json:"to_user_id,omitempty"`
	IsFollow bool  `json:"is_follow,omitempty"`
}

func (v Attention) TableName() string {
	return "attention"
}

type RelationDao struct {
}

var relationDao *RelationDao
var relationOnce sync.Once

// NewRelationDaoInstance Singleton
func NewRelationDaoInstance() *RelationDao {
	relationOnce.Do(
		func() {
			relationDao = &RelationDao{}
		})
	return relationDao
}

// QueryFollowList 查询关注者的id、姓名、关注数、粉丝数、与当前用户的关系
func (s *RelationDao) QueryFollowList(id int64) ([]*User, error) {
	var follows []Attention
	err := db.Debug().Table("attention").Where("user_id", id).Find(&follows).Error
	fmt.Println("关注列表", follows)
	var users []*User
	for _, follow := range follows {
		var user *User
		err = db.Table("users").Where("id = ?", follow.ToUserId).First(&user).Error
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, err
}

// QueryFollowerList 查询粉丝的id、姓名、关注数、粉丝数、与当前用户的关系
func (s *RelationDao) QueryFollowerList(id int64) ([]*User, error) {
	var follows []Attention
	err := db.Debug().Table("attention").Where("to_user_id", id).Find(&follows).Error
	fmt.Println("粉丝列表", follows)
	var users []*User
	for _, follow := range follows {
		var user *User
		err = db.Table("users").Where("id = ?", follow.UserId).First(&user).Error
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, err
}

// IsFollow 判断是否关注
func (s *RelationDao) IsFollow(userID int64, toUserId int64) bool {
	var att Attention
	if err := db.Table("attention").Select("*").Where("user_id = ? and to_user_id =?", userID, toUserId).First(&att).Error; err != nil {
		fmt.Println(err)
	}
	fmt.Println(att)
	return att.IsFollow
}

// FollowAction 关注操作
func (s *RelationDao) FollowAction(userId, toUserId int64) error {
	var att Attention
	result := db.Where("user_id = ? And to_user_id = ?", userId, toUserId).Find(&att)
	if result.RowsAffected != 0 {
		return utils.Error{
			Msg: "当前用户您已关注，无需重复关注! ",
		}
	}
	//对关系表的操作
	relation := Attention{
		UserId:   userId,
		ToUserId: toUserId,
		IsFollow: true,
	}
	db.Debug().Create(&relation)
	//对用户表的操作
	db.Debug().Where("id = ?", userId).First(&entity.User{}).Update("follow_count", gorm.Expr("follow_count + ?", 1))
	db.Debug().Where("id = ?", toUserId).First(&entity.User{}).Update("follower_count", gorm.Expr("follower_count + ?", 1))
	return nil
}

// WithdrawFollowAction 取关操作
func (s *RelationDao) WithdrawFollowAction(userId, toUserId int64) error {
	//对关系表的操作
	if err := db.Debug().Where("user_id=? and to_user_id =?", userId, toUserId).Delete(&Attention{}).Error; err != nil {
		return err
	}
	//对用户表的操作
	db.Debug().Where("id = ?", userId).First(&entity.User{}).Update("follow_count", gorm.Expr("follow_count - ?", 1))
	db.Debug().Where("id = ?", toUserId).First(&entity.User{}).Update("follower_count", gorm.Expr("follower_count - ?", 1))
	return nil
}
