package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/thisIsCaiJi/online_college/service_edu/model"
	"github.com/thisIsCaiJi/online_college/service_edu/util"
)

func init(){
	group.POST("/eduservice/course/addCourseInfo",AddCourseInfo)
}

func AddCourseInfo(ctx *gin.Context){
	courseVo := &model.EduCourseVo{}
	course := &model.EduCourse{}
	courseDescription := &model.EduCourseDescription{}
	ctx.ShouldBindJSON(courseVo)
	util.CopyStruct(courseVo,course)
	util.CopyStruct(courseVo,courseDescription)

	if err := course.CreateAll(courseDescription);err != nil {
		HandlerError(err,"添加课程信息失败")
		JsonErrorMessage(ctx,"添加课程信息失败")
		return
	}
	JsonSuccessKV(ctx,"id",course.Id)
}