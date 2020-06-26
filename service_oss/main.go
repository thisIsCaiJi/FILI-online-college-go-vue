package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thisIsCaiJi/online_college/service_oss/router"
	"github.com/thisIsCaiJi/online_college/service_oss/util"

	"log"
)

func main() {
	// Engin
	app := gin.Default()
	//router := gin.New()

	app.GET("/", func(context *gin.Context) {
		log.Println(">>>> hello gin start <<<<")
		context.JSON(200,util.ReturnOk().H())
	})

	router.Router(app)
	// 指定地址和端口号
	app.Run(":8082")
	//ossService := service.NewOssService()
	//url := ossService.UploadFileAvtar()
	//fmt.Printf("%s\n",url)
}
