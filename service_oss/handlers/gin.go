package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/thisIsCaiJi/online_college/service_oss/config"
	"github.com/thisIsCaiJi/online_college/service_oss/util"
)

var app = gin.Default()

func init(){
	app.Use(cors)
}

func cors(ctx *gin.Context){
	ctx.Header("Access-Control-Allow-Origin", "*")             //允许访问所有域
	ctx.Header("Access-Control-Allow-Headers", "Content-Type,x-token") //header的类型
	ctx.Header("content-type", "application/json")             //返回数据格式是json
	ctx.Header("Access-Control-Allow-Methods","GET,POST,PUT,PATCH,DELETE,OPTIONS")
	if ctx.Request.Method == "OPTIONS" {
		ctx.JSON(200,util.ReturnOk().H())
		return
	}
	ctx.Next()
}

func ServerRun() {
	port := config.Yml.ServerConf.Port
	app.Run(fmt.Sprintf(":%d",port))
}

func Close(){
	//model.Close()
}
