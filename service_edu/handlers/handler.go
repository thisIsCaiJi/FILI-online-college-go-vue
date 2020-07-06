package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func printError(err error, message string, a ...interface{}) {
	logrus.Errorf("ERROR: message: %s, err: %v\n", fmt.Sprintf(message, a...), err)
}

func handleError(c *gin.Context, err error) bool {
	if err != nil {
		jsonErrorMessage(c, err.Error())
		return true
	}
	return false
}

//返回成功
func jsonSuccess(ctx *gin.Context) {
	ctx.JSON(200, ReturnOk().H())
}

//返回失败
func jsonError(ctx *gin.Context) {
	ctx.JSON(200, ReturnError().H())
}


//返回带信息的失败
func jsonErrorMessage(ctx *gin.Context, message string) {
	ctx.JSON(200, ReturnError().Message(message).H())
}

//返回带数据的成功
func jsonSuccessKV(ctx *gin.Context, key string, value interface{}) {
	ctx.JSON(200, ReturnOk().DataKV(key, value).H())
}

func jsonSuccessMap(ctx *gin.Context, m map[string]interface{}) {
	ctx.JSON(200, ReturnOk().Data(m).H())
}

const SUCCESS_CODE = 20000

const FAILED_CODE = 20001

type Result struct {
	code int
	success bool
	message string
	data map[string]interface{}
}

func ReturnOk() Result {
	result := Result{
		code:SUCCESS_CODE,
		success:true,
		message:"成功",
		data:make(map[string]interface{},0),
	}
	return result
}

func ReturnError() Result {
	result := Result{
		code:FAILED_CODE,
		success:false,
		message:"失败",
		data:make(map[string]interface{},0),
	}
	return result
}

func (r Result) H() gin.H {
	return gin.H{
		"code":    r.code,
		"success": r.success,
		"message": r.message,
		"data":    r.data,
	}
}

func (r Result) Success(success bool) Result {
	r.success = success
	return r
}

func (r Result) Code(code int) Result {
	r.code = code
	return r
}

func (r Result) Message(message string) Result {
	r.message = message
	return r
}

func (r Result) Data(data map[string]interface{}) Result {
	r.data = data
	return r
}

func (r Result) DataKV(key string,value interface{}) Result {
	r.data[key] = value
	return r
}
