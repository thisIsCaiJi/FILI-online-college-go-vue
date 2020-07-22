package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/thisIsCaiJi/online_college/service_oss/util"
)


func UploadFileAvtar(context *gin.Context)  {
	file, header, err := context.Request.FormFile("file")
	if handleError(context,err){
		return
	}
	dataLen := header.Size
	if dataLen > int64(2)*1024*1024 {
		jsonErrorMessage(context,"File size cannot be greater than 2MB!")
		return
	}
	// 上传文件流。
	url,err := util.UploadFile(file,header)
	if handleError(context,err){
		return
	}
	jsonSuccessKV(context,"url",url)
}
