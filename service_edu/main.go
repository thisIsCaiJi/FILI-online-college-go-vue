package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thisIsCaiJi/online_college/service_edu/model"
	"github.com/thisIsCaiJi/online_college/service_edu/router"
	"github.com/thisIsCaiJi/online_college/service_edu/util"
	"log"
)

func main() {
	//初始化数据库
	model.DB.Init()
	defer model.DB.Close()

	app := gin.Default()

	app.GET("/", func(context *gin.Context) {
		log.Println(">>>> hello gin start <<<<")
		context.JSON(200,util.ReturnOk().H())
	})

	router.Router(app)
	// 指定地址和端口号
	app.Run(":8081")
}
