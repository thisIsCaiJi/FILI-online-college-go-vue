package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/thisIsCaiJi/online_college/service_edu/model"
	"github.com/thisIsCaiJi/online_college/service_edu/util"
	"strconv"
)

func init(){
	group := app.Group("/eduservice/edu-teacher")
	group.POST("/pageTeacherCondition/:current/:limit",PageTeacherConditionPost)

	group.POST("/addTeacher",AddTeacher)

	group.GET("/getTeacher/:id",GetTeacherById)

	group.POST("/updateTeacher",UpdateTeacher)

	group.DELETE("/:id",RemoveTeacher)
}

func PageTeacherConditionPost(ctx *gin.Context) {
	var query model.TeacherQuery
	var eduteacher model.EduTeacher
	err := ctx.ShouldBindJSON(&query)
	if err != nil {
		logrus.Infof("获取参数失败，err:%v\n", err)
	}
	current, limit := ctx.Param("current"), ctx.Param("limit")
	var currentInt,limitInt int
	if currentInt,err = strconv.Atoi(current);err != nil {
		currentInt = 1
	}
	if limitInt,err = strconv.Atoi(limit);err != nil {
		limitInt = 5
	}
	teachers, total, err := eduteacher.List(currentInt, limitInt, query)
	if err != nil {
		logrus.Errorf("分页查询讲师数据失败,err:%v\n", err)
		ctx.JSON(200, util.ReturnError().H())
		return
	}
	ctx.JSON(200, util.ReturnOk().DataKV("rows", teachers).DataKV("total", total).H())
}

func AddTeacher(ctx *gin.Context) {
	var eduTeacher model.EduTeacher
	ctx.ShouldBindJSON(&eduTeacher)
	err := eduTeacher.Add()
	if err != nil {
		logrus.Errorf("添加讲师数据失败,err:%v\n", err)
		ctx.JSON(200, util.ReturnError().H())
		return
	}
	ctx.JSON(200, util.ReturnOk().H())
}

func GetTeacherById(ctx *gin.Context) {
	var eduTeacher model.EduTeacher
	id := ctx.Param("id")
	eduTeacher.Id = id
	teacher, err := eduTeacher.GetById()
	if err != nil {
		logrus.Errorf("查询讲师数据失败,err:%v\n", err)
		ctx.JSON(200, util.ReturnError().H())
		return
	}
	ctx.JSON(200, util.ReturnOk().DataKV("teacher",teacher).H())
}

func UpdateTeacher(ctx *gin.Context) {
	var eduTeacher model.EduTeacher
	ctx.ShouldBindJSON(&eduTeacher)
	t,err := eduTeacher.GetById()
	if err != nil {
		logrus.Errorf("修改讲师数据失败,err:%v\n", err)
		ctx.JSON(200, util.ReturnError().H())
		return
	}
	eduTeacher.GmtCreate = t.GmtCreate
	err = eduTeacher.Update()
	if err != nil {
		logrus.Errorf("修改讲师数据失败,err:%v\n", err)
		ctx.JSON(200, util.ReturnError().H())
		return
	}
	ctx.JSON(200, util.ReturnOk().H())
}

func RemoveTeacher(ctx *gin.Context) {
	id := ctx.Param("id")
	eduTeacher := model.EduTeacher{Id:id}
	err := eduTeacher.Remove()
	if err != nil {
		logrus.Errorf("逻辑删除讲师数据失败,err:%v\n", err)
		ctx.JSON(200, util.ReturnError().H())
		return
	}
	ctx.JSON(200, util.ReturnOk().H())
}
