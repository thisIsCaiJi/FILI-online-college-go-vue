package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func HandlerError(err error, message string, a ...interface{}) {
	logrus.Errorf("ERROR: message: %s, err: %v\n", fmt.Sprintf(message, a...), err)
}

//返回成功
func JsonSuccess(ctx *gin.Context) {
	ctx.JSON(200, ReturnOk().H())
}

//返回失败
func JsonError(ctx *gin.Context) {
	ctx.JSON(200, ReturnError().H())
}

//返回带信息的失败
func JsonErrorMessage(ctx *gin.Context, message string) {
	ctx.JSON(200, ReturnError().Message(message).H())
}

//返回带数据的成功
func JsonSuccessKV(ctx *gin.Context, key string, value interface{}) {
	ctx.JSON(200, ReturnOk().DataKV(key, value).H())
}

func JsonSuccessMap(ctx *gin.Context, m map[string]interface{}) {
	ctx.JSON(200, ReturnOk().Data(m).H())
}
