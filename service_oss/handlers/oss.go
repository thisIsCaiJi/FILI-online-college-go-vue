package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/thisIsCaiJi/online_college/service_oss/util"
)

func init(){
	group := app.Group("/eduoss")
	group.POST("/fileoss",UploadFileAvtar)
}

func UploadFileAvtar(context *gin.Context)  {
	file, header, err := context.Request.FormFile("file")
	if err != nil {
		fmt.Printf("上传失败")
		context.JSON(200,util.ReturnError().H())
		return
	}
	dataLen := header.Size
	if dataLen > int64(2)*1024*1024 {
		context.JSON(200,util.ReturnError().Message("File size cannot be greater than 2MB!").H())
		return
	}
	// 上传文件流。
	url,err := util.UploadFile(file,header)
	if err != nil {
		context.JSON(200,util.ReturnError().Message(err.Error()).H())
		return
	}
	context.JSON(200,util.ReturnOk().DataKV("url",url).H())
}
