package Impl

import "github.com/Usigned/douyin/entity"

// UserRepository ..
type UserRepository interface {
	Save(user *entity.User) error
	TotalNum() (int64, error)
	// FindByID(ID int) (*common.User, error)
	FindByName(Name string) (*entity.User, error)
}
