package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/thisIsCaiJi/online_college/service_edu/model"
	"strconv"
)



func PageTeacherConditionPost(ctx *gin.Context) {
	var query model.TeacherQuery
	var eduteacher model.EduTeacher
	err := ctx.ShouldBindJSON(&query)
	if err != nil {
		logrus.Infof("获取参数失败，err:%v\n", err)
	}
	current, limit := ctx.Param("current"), ctx.Param("limit")
	var currentInt, limitInt int
	if currentInt, err = strconv.Atoi(current); err != nil {
		currentInt = 1
	}
	if limitInt, err = strconv.Atoi(limit); err != nil {
		limitInt = 5
	}
	teachers, total, err := eduteacher.List(currentInt, limitInt, query)
	if err != nil {
		logrus.Errorf("分页查询讲师数据失败,err:%v\n", err)
		jsonError(ctx)
		return
	}
	jsonSuccessMap(ctx, map[string]interface{}{"rows": teachers, "total": total})
}

func AddTeacher(ctx *gin.Context) {
	var eduTeacher model.EduTeacher
	ctx.ShouldBindJSON(&eduTeacher)
	eduTeacher.Id = ""
	err := eduTeacher.Add()
	if err != nil {
		logrus.Errorf("添加讲师数据失败,err:%v\n", err)
		jsonError(ctx)
		return
	}
	jsonSuccess(ctx)
}

func GetTeacherById(ctx *gin.Context) {
	var eduTeacher model.EduTeacher
	id := ctx.Param("id")
	eduTeacher.Id = id
	teacher, err := eduTeacher.GetById()
	if err != nil {
		logrus.Errorf("查询讲师数据失败,err:%v\n", err)
		jsonError(ctx)
		return
	}
	jsonSuccessKV(ctx, "teacher", teacher)
}

func UpdateTeacher(ctx *gin.Context) {
	var eduTeacher model.EduTeacher
	ctx.ShouldBindJSON(&eduTeacher)
	if eduTeacher.Id == "" {
		jsonErrorMessage(ctx,"id不能为空")
		return
	}
	t, err := eduTeacher.GetById()
	if err != nil {
		logrus.Errorf("修改讲师数据失败,err:%v\n", err)
		jsonSuccess(ctx)
		return
	}
	eduTeacher.GmtCreate = t.GmtCreate
	err = eduTeacher.Update()
	if err != nil {
		logrus.Errorf("修改讲师数据失败,err:%v\n", err)
		jsonError(ctx)
		return
	}
	jsonSuccess(ctx)
}

func RemoveTeacher(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		jsonErrorMessage(ctx,"id不能为空")
		return
	}
	eduTeacher := model.EduTeacher{Id: id}
	err := eduTeacher.Remove()
	if err != nil {
		logrus.Errorf("逻辑删除讲师数据失败,err:%v\n", err)
		jsonError(ctx)
		return
	}
	jsonSuccess(ctx)
}

func TeacherList(ctx *gin.Context) {
	eduTeacher := &model.EduTeacher{}
	teachers, total, err := eduTeacher.All()
	if err != nil {
		printError(err, "查询讲师数据失败")
		jsonErrorMessage(ctx, "查询讲师数据失败")
		return
	}
	jsonSuccessMap(ctx, map[string]interface{}{"rows": teachers, "total": total})
}
