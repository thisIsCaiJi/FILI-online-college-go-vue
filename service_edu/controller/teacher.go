package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/thisIsCaiJi/online_college/service_edu/model"
	"github.com/thisIsCaiJi/online_college/service_edu/service"
	"github.com/thisIsCaiJi/online_college/service_edu/util"
)

func PageTeacherConditionPost(ctx *gin.Context) {
	var query model.TeacherQuery
	err := ctx.ShouldBindJSON(&query)
	if err != nil {
		logrus.Infof("获取参数失败，err:%v\n", err)
	}
	current, limit := ctx.Param("current"), ctx.Param("limit")
	teachers, total, err := service.PageTeacherConditionPost(current, limit, query)
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
	err := service.AddTeacher(eduTeacher)
	if err != nil {
		logrus.Errorf("添加讲师数据失败,err:%v\n", err)
		ctx.JSON(200, util.ReturnError().H())
		return
	}
	ctx.JSON(200, util.ReturnOk().H())
}

func GetTeacherById(ctx *gin.Context) {
	id := ctx.Param("id")
	teacher, err := service.GetTeacherById(id)
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
	err := service.UpdateTeacher(eduTeacher)
	if err != nil {
		logrus.Errorf("修改讲师数据失败,err:%v\n", err)
		ctx.JSON(200, util.ReturnError().H())
		return
	}
	ctx.JSON(200, util.ReturnOk().H())
}

func RemoveTeacher(ctx *gin.Context) {
	id := ctx.Param("id")
	err := service.RemoveTeacher(id)
	if err != nil {
		logrus.Errorf("逻辑删除讲师数据失败,err:%v\n", err)
		ctx.JSON(200, util.ReturnError().H())
		return
	}
	ctx.JSON(200, util.ReturnOk().H())
}
