package entity

import "time"

type User struct {
	Id   int64
	Name string
}

type Video struct {
	Id       int64
	AuthorId int64
	PlayUrl  string
	CoverUrl string
	Title    string
	CreateAt time.Time
}
