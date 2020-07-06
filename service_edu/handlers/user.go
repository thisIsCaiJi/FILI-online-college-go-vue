package handlers

import (
	"github.com/gin-gonic/gin"
)


func Login(ctx *gin.Context) {
	jsonSuccessKV(ctx,"token","admin")
}

func Info(ctx *gin.Context) {
	jsonSuccessMap(ctx,map[string]interface{}{"roles": []string{"admin"},"name": "admin1","avatar":"https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif"})
}

func Logout(ctx *gin.Context) {
	jsonSuccess(ctx)
}
