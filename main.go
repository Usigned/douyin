package main

import (
	"github.com/Usigned/douyin/dao"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	initRouter(r)
	err := dao.Init(true)
	if err != nil {
		println(err.Error())
	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
