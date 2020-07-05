package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/thisIsCaiJi/online_college/service_edu/model"
	"github.com/thisIsCaiJi/online_college/service_edu/util"
)

func init(){
	group.POST("/eduservice/course/addCourseInfo",AddCourseInfo)
	group.GET("/eduservice/course/getCourseInfo/:id",GetCourseInfo)
	group.POST("/eduservice/course/updateCourseInfo",UpdateCourseInfo)
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

func GetCourseInfo(ctx *gin.Context){
	id := ctx.Param("id")
	course := &model.EduCourse{Id:id}
	courseVo,err := course.OneAll()
	if err != nil {
		JsonErrorMessage(ctx,"查询课程信息失败")
		return
	}
	JsonSuccessKV(ctx,"courseInfo",courseVo)
}

func UpdateCourseInfo(ctx *gin.Context){
	courseVo := &model.EduCourseVo{}
	ctx.ShouldBindJSON(courseVo)
	if courseVo.Id == "" {
		JsonErrorMessage(ctx,"课程id不能为空")
		return
	}
	course, description := &model.EduCourse{},&model.EduCourseDescription{Id:courseVo.Id,Description:courseVo.Description}
	util.CopyStruct(courseVo,course)
	if err := course.UpdateAll(description);err!=nil {
		JsonErrorMessage(ctx,"修改课程信息失败")
		return
	}
	JsonSuccess(ctx)
}