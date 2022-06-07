package main

import (
	"douyin/dao"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	initRouter(r)

	err := initDB()
	if err != nil {
		println(err.Error())
		return
	}
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func initDB() error {
	return dao.Init(true)
}
