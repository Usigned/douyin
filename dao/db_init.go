package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB
var dsn = "root:abc123@tcp(127.0.0.1:3306)/douyins?charset=utf8mb4&parseTime=True&loc=Local"

func Init(migrate bool) error {
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
		return err
	}

	// 自动迁移
	if migrate {
		err = Migrate()
		if err != nil {
			return err
		}
	}

	return nil
}

func Migrate() error {
	err := db.AutoMigrate(&User{}, &Comment{}, &Video{}, &Favorite{}, &LoginStatus{}, &Attention{})
	if err != nil {
		return err
	}
	return nil
}
