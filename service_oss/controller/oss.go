package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/thisIsCaiJi/online_college/service_oss/service"
	"github.com/thisIsCaiJi/online_college/service_oss/util"
)

func UploadFileAvtar(context *gin.Context)  {
	file, header, err := context.Request.FormFile("file")
	if err != nil {
		fmt.Printf("上传失败")
		context.JSON(200,util.ReturnError().H())
		return
	}
	url,err := service.UploadFileAvtar(file,header)
	if err != nil {
		context.JSON(200,util.ReturnError().Message(err.Error()).H())
		return
	}
	context.JSON(200,util.ReturnOk().DataKV("url",url).H())
}
