package router

import (
	"github.com/gin-gonic/gin"
	"github.com/thisIsCaiJi/online_college/service_oss/controller"
	"github.com/thisIsCaiJi/online_college/service_oss/util"
)

func Router(app *gin.Engine){
	group := app.Group("/eduoss",cors)
	group.POST("/fileoss",controller.UploadFileAvtar)
	group.OPTIONS("/fileoss",options)
}

func options(ctx *gin.Context) {
	ctx.JSON(200,util.ReturnOk().H())
}


func cors(ctx *gin.Context){
	ctx.Header("Access-Control-Allow-Origin", "*")             //允许访问所有域
	ctx.Header("Access-Control-Allow-Headers", "Content-Type,X-Token") //header的类型
	ctx.Header("content-type", "application/json")             //返回数据格式是json
	ctx.Next()
}