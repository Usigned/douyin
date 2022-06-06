package dao

import (
	"douyin/entity"
	"fmt"
	"gorm.io/gorm"
)

type Query struct {
	Token  string `form:"token"`
	UserId int64  `form:"user_id"`
}

type Change struct {
	Token      string `form:"token"`
	UserId     int64  `form:"user_id"`
	ToUserId   int64  `form:"to_user_id"`
	ActionType int64  `form:"action_type"`
}

type UserList struct {
	Id            int64  `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"`
}

type Att struct {
	UserId          int64 `json:"user_id,omitempty"`
	AttentionUserId int64 `json:"attention_user_id,omitempty"`
	IsFollow        bool  `json:"is_follow,omitempty"`
}

func (v Att) TableName() string {
	return "attention"
}

type AttentionDao struct {
}

// FindFollowList 查询关注者的id、姓名、关注数、粉丝数、与当前用户的关系
func FindFollowList(id int64) ([]UserList, error) {
	var followList []UserList
	if err := db.Debug().Table("users").Select("users.id,users.name,users.follow_count,users.follower_count,attention.is_follow").Joins("inner join attention on attention.attention_user_id=users.id").Where("attention.user_id=?", id).Find(&followList).Error; err != nil {
		return nil, err
	}
	fmt.Println("关注列表", followList)
	return followList, nil
}

// FindFollowerList 查询粉丝的id、姓名、关注数、粉丝数、与当前用户的关系
func FindFollowerList(id int64) ([]UserList, error) {
	var followerList []UserList
	if err := db.Debug().Table("users").Select("users.id,users.name,users.follow_count,users.follower_count,attention.is_follow").Joins("inner join attention on attention.user_id=users.id").Where("attention.attention_user_id=?", id).Find(&followerList).Error; err != nil {
		return nil, err
	}
	fmt.Println("粉丝列表", followerList)
	return followerList, nil
}

// FollowAction 关注操作
func FollowAction(idA int64, idB int64) error {
	//对关系表的操作
	//1.进行关注操作前，判断是否单向被关注，如果是，则将关系is_follow改成双向true，否则不做处理
	var reverse Att
	err := db.Debug().Where("user_id = ? and attention_user_id = ?", idB, idA).First(&reverse).Update("is_follow", true).Error
	fmt.Println(reverse)
	if err != nil {
		//如果对方没有关注自己，则进行插入时设置成单向
		fmt.Println("没查到：", err)
		relation := Att{
			UserId:          idA,
			AttentionUserId: idB,
			IsFollow:        false,
		}
		db.Debug().Create(&relation)
	} else {
		//如果对方已经关注自己，则进行插入时设置成双向
		fmt.Println("查到数据，并已经修改关系")
		relation := Att{
			UserId:          idA,
			AttentionUserId: idB,
			IsFollow:        true,
		}
		db.Debug().Create(&relation)
	}

	//对用户表的操作
	var userA entity.User
	//Update(“follow_count”, gorm.Expr(“follow_count + ?”, 1))
	db.Debug().Where("id = ?", idA).First(&userA).Update("follow_count", gorm.Expr("follow_count + ?", 1))
	var userB entity.User
	db.Debug().Where("id = ?", idB).First(&userB).Update("follower_count", gorm.Expr("follower_count + ?", 1))

	return nil
}

// FollowerAction 取关操作
func FollowerAction(idA int64, idB int64) error {
	//对关系表的操作

	//1.进行取关操作前，判断是否双向关注，如果是，则将关系is_follow改成单向false，否则不做处理
	var reverse Att
	db.Debug().Where("user_id = ? and attention_user_id = ?", idB, idA).First(&reverse).Update("is_follow", false)
	fmt.Println(reverse)
	//2.然后删除关注关系
	var relation Att
	if err := db.Debug().Where("user_id=? and attention_user_id =?", idA, idB).Delete(&relation).Error; err != nil {
		fmt.Println(relation)
		return err
	}
	//对用户表的操作
	//对用户表的操作
	var userA entity.User
	//Update(“follow_count”, gorm.Expr(“follow_count + ?”, 1))
	db.Debug().Where("id = ?", idA).First(&userA).Update("follow_count", gorm.Expr("follow_count - ?", 1))
	var userB entity.User
	db.Debug().Where("id = ?", idB).First(&userB).Update("follower_count", gorm.Expr("follower_count - ?", 1))
	return nil
}
