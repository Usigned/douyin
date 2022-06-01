package main

import (
	"github.com/Usigned/douyin/dao"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	initRouter(r)
	//err := dao.Init(true)
	//if err != nil {
	//	println(err.Error())
	//}

	err := initDB()
	if err != nil {
		println(err.Error())
		return
	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func initDB() error {
	return dao.Init(false)
}
