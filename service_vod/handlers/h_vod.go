package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/thisIsCaiJi/online_college/service_vod/service"
)

func UploadVideo(ctx *gin.Context) {
	file, header, err := ctx.Request.FormFile("file")
	if handleError(ctx,err) {
		return
	}
	dataLen := header.Size
	if dataLen > int64(10)*1024*1024 {
		jsonErrorMessage(ctx,"File size cannot be greater than 10MB!")
		return
	}
	videoId, err := service.UploadVideo(file,header)
	if handleError(ctx,err) {
		return
	}
	jsonSuccessKV(ctx,"videoId",videoId)
}

func GetVideo(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		jsonErrorMessage(ctx,"id不能为空")
		return
	}
}

func DeleteVideo(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		jsonErrorMessage(ctx,"id不能为空")
		return
	}
	err := service.DeleteVideo(id)
	if handleError(ctx,err) {
		return
	}
	jsonSuccess(ctx)
}
