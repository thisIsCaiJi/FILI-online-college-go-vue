package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/thisIsCaiJi/online_college/service_edu/util"
)

func Login(ctx *gin.Context) {
	ctx.JSON(200, util.ReturnOk().DataKV("token", "admin").H())
}

func Info(ctx *gin.Context) {
	ctx.JSON(200, util.ReturnOk().DataKV("roles", []string{"admin"}).DataKV("name", "admin1").
		DataKV("avatar", "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif").H())
}

func Logout(ctx *gin.Context) {
	ctx.JSON(200, util.ReturnOk().H())
}
