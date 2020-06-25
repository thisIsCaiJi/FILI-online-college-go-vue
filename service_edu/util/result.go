package util

import (
	"github.com/gin-gonic/gin"
)

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