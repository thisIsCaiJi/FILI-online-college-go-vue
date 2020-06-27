package model

import (
	"github.com/sirupsen/logrus"
	"github.com/thisIsCaiJi/online_college/service_edu/util"
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

func (*EduTeacher) TableName() string {
	return "edu_teacher"
}

type TeacherQuery struct {
	Name  string `json:"name"`
	Level int    `json:"level"`
	Begin string `json:"begin"`
	End   string `json:"end"`
}

func (t *EduTeacher) List(current, limit int,query TeacherQuery) (*[]EduTeacher,int,error) {
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
	if err := db.Where(whereSql,whereValue...).Limit(limit).Offset((current-1) * limit).Find(&eduTeacher).Error;err != nil {
		return nil,0,err
	}
	if err := db.Model(t).Where(whereSql,whereValue...).Count(&total).Error;err != nil {
		return nil,0,err
	}
	return &eduTeacher,total,nil
}

func (t *EduTeacher) Add() error {
	id,err := util.GetId()
	if err!=nil {
		return err
	}
	t.Id = id
	t.GmtCreate = time.Now()
	t.GmtModified = time.Now()
	if err := db.Create(t).Error; err != nil {
		return err
	}
	return nil
}

func (t *EduTeacher) GetById() (EduTeacher,error) {
	var teacher EduTeacher
	err := db.Where("id = ?",t.Id).Find(&teacher).Error
	return teacher,err
}

func (t *EduTeacher) Update() error{
	t.GmtModified = time.Now()
	err := db.Save(t).Error
	return err
}

func (t *EduTeacher) Remove() error{
	err := db.Model(t).Where("id = ?",t.Id).Update("is_deleted",1).Error
	return err
}
