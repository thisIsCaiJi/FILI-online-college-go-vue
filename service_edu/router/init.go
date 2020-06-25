package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/thisIsCaiJi/online_college/service_edu/util"
)

func Router(app *gin.Engine){
	app.Use(cors)
	uGroup := app.Group("/eduservice/user")
	userRouter(uGroup)
	teacherGroup := app.Group("/eduservice/edu-teacher")
	teacherRouter(teacherGroup)
}

func cors(ctx *gin.Context){
	fmt.Printf("request mode:%v\n",ctx.Request.Method)
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