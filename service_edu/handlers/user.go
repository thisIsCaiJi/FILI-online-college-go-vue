package handlers

import (
	"github.com/gin-gonic/gin"
)

func init() {
	group.POST("/eduservice/user/login",Login)
	group.GET("/eduservice/user/info",Info)
	group.POST("/eduservice/user/logout",Logout)
}

func Login(ctx *gin.Context) {
	JsonSuccessKV(ctx,"token","admin")
}

func Info(ctx *gin.Context) {
	JsonSuccessMap(ctx,map[string]interface{}{"roles":[]string{"admin"},"name": "admin1","avatar":"https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif"})
}

func Logout(ctx *gin.Context) {
	JsonSuccess(ctx)
}
