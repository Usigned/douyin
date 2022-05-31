package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
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
