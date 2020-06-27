package handlers

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gin-gonic/gin"
	"github.com/thisIsCaiJi/online_college/service_edu/model"
	"github.com/thisIsCaiJi/online_college/service_edu/util"
	"path"
)

func init(){
	group := app.Group("/eduservice/subject",cors)

	group.POST("/import",ImportSubject)
	group.GET("/allSubject",GetAllSubject)
}

func ImportSubject(ctx *gin.Context) {
	file, header, err := ctx.Request.FormFile("file")
	if err != nil {
		HandlerError(err,"上传失败")
		ctx.JSON(200,util.ReturnError().H())
		return
	}
	ty := path.Ext(header.Filename)
	if ty != ".xlsx" {
		ctx.JSON(200,util.ReturnError().Message("只能上传.xlsx表格文件").H())
		return
	}
	xlFile,err := excelize.OpenReader(file)
	if err != nil {
		HandlerError(err,"上传失败")
		ctx.JSON(200,util.ReturnError().Message("上传失败").H())
		return
	}
	rows := xlFile.GetRows("课程信息")
	for irow, row := range rows {
		if irow > 0 {
			if len(row) >= 2 {
				subjectSave := model.EduSubject{Title:row[0],ParentId:"0"}
				subjectOld,notexist := subjectSave.GetByTitleAndParent()
				var id string
				if notexist {
					err = subjectSave.Add()
					if err != nil {
						HandlerError(err,"导入一级课程错误")
						continue
					}
					id = subjectSave.Id
				}else {
					id = subjectOld.Id
				}
				subjectSave = model.EduSubject{Title:row[1],ParentId:id}
				subjectOld,notexist = subjectSave.GetByTitleAndParent()
				if notexist {
					err = subjectSave.Add()
					if err != nil {
						HandlerError(err,"导入二级课程错误")
						continue
					}
				}
			}
		}
	}
	ctx.JSON(200,util.ReturnOk().H())
}

//课程分类列表，树形结构
func GetAllSubject(ctx *gin.Context){
	var subject model.EduSubject
	subject.ParentId = "0"
	oneSub,_,err := subject.List()
	if err != nil {
		HandlerError(err,"查询一级课程错误")
		return
	}
	twoSub,err := subject.GetTwoSub()
	if err != nil {
		HandlerError(err,"查询二级课程错误")
		return
	}
	list := make([]model.NodeSubject,0,len(*oneSub))
	m := make(map[string]*[]model.NodeSubject)
	for _,sub := range *oneSub {
		one := model.NodeSubject{
			Id:       sub.Id,
			Title:    sub.Title,
			Children: &[]model.NodeSubject{},
		}
		list = append(list,one)
		m[one.Id] = one.Children
	}
	for _,sub2 := range *twoSub {
		if _,ok := m[sub2.ParentId]; ok{
			two := model.NodeSubject{
				Id:       sub2.Id,
				Title:    sub2.Title,
			}
			*m[sub2.ParentId] = append(*m[sub2.ParentId],two)
		}
	}
	ctx.JSON(200,util.ReturnOk().DataKV("list",list).H())
}

