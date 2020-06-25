package router

import (
	"github.com/gin-gonic/gin"
	"github.com/thisIsCaiJi/online_college/service_edu/controller"
)

func teacherRouter(group *gin.RouterGroup){
	group.POST("/pageTeacherCondition/:current/:limit",controller.PageTeacherConditionPost)

	group.POST("/addTeacher",controller.AddTeacher)

	group.GET("/getTeacher/:id",controller.GetTeacherById)

	group.POST("/updateTeacher",controller.UpdateTeacher)

	group.DELETE("/:id",controller.RemoveTeacher)
}
