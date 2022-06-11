package main

import (
	"douyin/dao/mysql"
	"douyin/logger"
	"douyin/settings"
	"fmt"
	"github.com/gin-gonic/gin"
)

type any = interface{}

func main() {
	r := gin.Default()

	initRouter(r)

	// 加载配置
	if err := settings.Init(); err != nil {
		fmt.Printf("load config failed, err:%v\n", err)
		return
	}
	if err := logger.Init(settings.Conf.LogConfig, settings.Conf.Mode); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}
	if err := mysql.Init(settings.Conf.MySQLConfig, false); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}

	//if err := redis.Init(settings.Conf.RedisConfig); err != nil {
	//	fmt.Printf("init redis failed, err:%v\n", err)
	//	return
	//}
	//defer redis.Close()

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
