package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

var db *gorm.DB

func Init() error {
	var err error
	dsn := "qing:cdma1330@tcp(118.202.10.86:3306)/douyin?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
		return err
	}
	err = Migrate()
	if err != nil {
		return err
	}
	return nil
}

func Migrate() error {
	err := db.AutoMigrate(&Video{})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&User{})
	if err != nil {
		return err

	}
	return nil
}

func InsertDemoUser() error {
	var user = &User{
		Name:     "qing",
		Password: "123456",
	}
	return db.Create(user).Error
}

func InsertDemoVideo() error {
	var video = &Video{
		AuthorId: 1,
		PlayUrl:  "https://www.w3schools.com/html/movie.mp4",
		CoverUrl: "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		Title:    "Bear",
		CreateAt: time.Now(),
	}
	return db.Create(video).Error
}
