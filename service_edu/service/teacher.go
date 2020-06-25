package service

import (
	"github.com/thisIsCaiJi/online_college/service_edu/model"
	"strconv"
)

func PageTeacherConditionPost(current, limit string,query model.TeacherQuery) ([]model.EduTeacher, int, error) {
	currentInt,err := strconv.Atoi(current)
	if err != nil {
		currentInt = 1
	}
	limitInt,err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 5
	}
	teachers, total := model.TeacherList(currentInt, limitInt, query)
	return teachers,total,nil
}

func AddTeacher(teacher model.EduTeacher) error{
	return model.AddTeacher(teacher)
}

func GetTeacherById(id string) (model.EduTeacher,error){
	teacher,err := model.GetTeacherById(id)
	return teacher,err
}

func UpdateTeacher(teacher model.EduTeacher) error{
	teacherOld,err := GetTeacherById(teacher.Id)
	if err != nil {
		return err
	}
	teacher.GmtCreate = teacherOld.GmtCreate
	err = model.UpdateTeacher(teacher)
	return err
}

func RemoveTeacher(id string) error{
	return model.RemoveTeacher(id)
}
