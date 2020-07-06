package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/thisIsCaiJi/online_college/service_edu/model"
	"github.com/thisIsCaiJi/online_college/service_edu/util"
)


func AddCourseInfo(ctx *gin.Context){
	courseVo := &model.EduCourseVo{}
	course := &model.EduCourse{}
	courseDescription := &model.EduCourseDescription{}
	ctx.ShouldBindJSON(courseVo)
	util.CopyStruct(courseVo,course)
	util.CopyStruct(courseVo,courseDescription)

	if err := course.CreateAll(courseDescription);err != nil {
		printError(err,"添加课程信息失败")
		jsonErrorMessage(ctx,"添加课程信息失败")
		return
	}
	jsonSuccessKV(ctx,"id",course.Id)
}

func GetCourseInfo(ctx *gin.Context){
	id := ctx.Param("id")
	course := &model.EduCourse{Id:id}
	courseVo,err := course.OneAll()
	if err != nil {
		jsonErrorMessage(ctx,"查询课程信息失败")
		return
	}
	jsonSuccessKV(ctx,"courseInfo",courseVo)
}

func UpdateCourseInfo(ctx *gin.Context){
	courseVo := &model.EduCourseVo{}
	ctx.ShouldBindJSON(courseVo)
	if courseVo.Id == "" {
		jsonErrorMessage(ctx,"课程id不能为空")
		return
	}
	course, description := &model.EduCourse{},&model.EduCourseDescription{Id:courseVo.Id,Description:courseVo.Description}
	util.CopyStruct(courseVo,course)
	if err := course.UpdateAll(description);err!=nil {
		jsonErrorMessage(ctx,"修改课程信息失败")
		return
	}
	jsonSuccess(ctx)
}

func GetCoursePublishVo(ctx *gin.Context){
	id := ctx.Param("id")
	if id == ""{
		jsonError(ctx)
		return
	}
	course := &model.EduCourse{Id:id}
	coursePublish, err := course.GetCoursePublishVo()
	if handleError(ctx,err){
		return
	}
	jsonSuccessKV(ctx,"coursePublish",coursePublish)
}