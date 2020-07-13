package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/thisIsCaiJi/online_college/service_edu/model"
)

func AddVideo(ctx *gin.Context){
	video := &model.EduVideo{}
	ctx.ShouldBindJSON(video)
	video.Id = ""
	err := video.Create()
	if handleError(ctx,err) {
		return
	}
	jsonSuccess(ctx)
}

func GetVideo(ctx *gin.Context){
	id := ctx.Param("id")
	videoParam := &model.EduVideo{Id:id}
	video,err := videoParam.One()
	if handleError(ctx,err) {
		return
	}
	jsonSuccessKV(ctx,"video",video)
}

func RemoveVideo(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		jsonErrorMessage(ctx,"id不能为空")
		return
	}
	videoParam := &model.EduVideo{Id:id}
	err := videoParam.Delete()
	if handleError(ctx,err) {
			return
		}
	jsonSuccess(ctx)
}

func UpdateVideo(ctx *gin.Context) {
	video := &model.EduVideo{}
	ctx.ShouldBindJSON(video)
	if video.Id == "" {
		jsonErrorMessage(ctx,"id不能为空")
		return
	}
	err := video.Update()
	if handleError(ctx,err) {
		return
	}
	jsonSuccess(ctx)
}