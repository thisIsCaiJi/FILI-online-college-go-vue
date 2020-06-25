package model

import (
	"fmt"
	"github.com/gitstliu/go-id-worker"
	"github.com/sirupsen/logrus"
	"strconv"
	"time"
)

type EduTeacher struct {
	Id          string     `json:"id" gorm:"primary_key"`
	Name        string     `json:"name"`
	Intro       string     `json:"intro"`
	Career      string     `json:"career"`
	Level       int        `json:"level"`
	Avatar      string     `json:"avatar"`
	Sort        int        `json:"sort"`
	IsDeleted   int        `json:"isDeleted"`
	GmtCreate   time.Time `json:"gmtCreate"`
	GmtModified time.Time `json:"gmtModified"`
}

func (EduTeacher) TableName() string {
	return "edu_teacher"
}

type TeacherQuery struct {
	Name  string `json:"name"`
	Level int    `json:"level"`
	Begin string `json:"begin"`
	End   string `json:"end"`
}

func TeacherList(current, limit int,query TeacherQuery) ([]EduTeacher,int) {
	logrus.Infof("query:%v\n",query)
	var eduTeacher []EduTeacher
	var total int
	whereSql := "is_deleted = 0"
	var whereValue []interface{}
	if query.Name != "" {
		whereSql += " and name like ?"
		whereValue = append(whereValue,"%"+query.Name+"%")
	}
	if query.Level != 0 {
		whereSql += " and level = ?"
		whereValue = append(whereValue,query.Level)
	}
	if query.Begin != "" {
		whereSql += " and gmt_create >= ?"
		whereValue = append(whereValue,query.Begin)
	}
	if query.End != "" {
		whereSql += " and gmt_create < ?"
		whereValue = append(whereValue,query.End)
	}
	DB.Self.Where(whereSql,whereValue...).Limit(limit).Offset((current-1) * limit).Find(&eduTeacher)
	DB.Self.Table(EduTeacher{}.TableName()).Where(whereSql,whereValue...).Count(&total)
	fmt.Printf("total:%v\n",total)
	return eduTeacher,total
}

func AddTeacher(teacher EduTeacher) error {
	currWoker := &idworker.IdWorker{}
	currWoker.InitIdWorker(1000, 1)
	id,err := currWoker.NextId()
	if err!=nil {
		return err
	}
	teacher.Id = strconv.FormatInt(id,10)
	teacher.GmtCreate = time.Now()
	teacher.GmtModified = time.Now()
	if err := DB.Self.Create(teacher).Error; err != nil {
		return err
	}
	return nil
}

func GetTeacherById(id string) (EduTeacher,error) {
	var teacher EduTeacher
	err := DB.Self.Where("id = ?",id).Find(&teacher).Error
	return teacher,err
}

func UpdateTeacher(teacher EduTeacher) error{
	teacher.GmtModified = time.Now()
	err := DB.Self.Save(&teacher).Error
	return err
}

func RemoveTeacher(id string) error{
	err := DB.Self.Model(&EduTeacher{}).Where("id = ?",id).Update("is_deleted",1).Error
	return err
}
