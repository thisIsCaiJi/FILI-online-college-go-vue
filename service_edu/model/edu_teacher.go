package model

import (
	"github.com/sirupsen/logrus"
	"time"
)

type EduTeacher struct {
	Id          string    `json:"id" gorm:"primary_key"`
	Name        string    `json:"name"`
	Intro       string    `json:"intro"`
	Career      string    `json:"career"`
	Level       int       `json:"level"`
	Avatar      string    `json:"avatar"`
	Sort        int       `json:"sort"`
	IsDelete    int `json:"is_delete"`
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

func (t *EduTeacher) List(current, limit int, query TeacherQuery) (*[]EduTeacher, int, error) {
	logrus.Infof("query:%v\n", query)
	var eduTeacher []EduTeacher
	var total int
	whereSql := "is_delete = 0"
	var whereValue []interface{}
	if query.Name != "" {
		whereSql += " and name like ?"
		whereValue = append(whereValue, "%"+query.Name+"%")
	}
	if query.Level != 0 {
		whereSql += " and level = ?"
		whereValue = append(whereValue, query.Level)
	}
	if query.Begin != "" {
		whereSql += " and gmt_create >= ?"
		whereValue = append(whereValue, query.Begin)
	}
	if query.End != "" {
		whereSql += " and gmt_create < ?"
		whereValue = append(whereValue, query.End)
	}
	if err := db.Model(t).Where(whereSql, whereValue...).Count(&total).Limit(limit).Offset((current - 1) * limit).Find(&eduTeacher).Error; err != nil {
		return nil, 0, err
	}
	return &eduTeacher, total, nil
}

func (t *EduTeacher) Add() error {
	return Create(t)
}

func (t *EduTeacher) GetById() (one *EduTeacher,err error) {
	one = &EduTeacher{}
	m := &EduTeacher{Id:t.Id,IsDelete:0}
	err = One(m,one)
	return
}

func (t *EduTeacher) Update() error {
	err := Update(t)
	return err
}

func (t *EduTeacher) Remove() error {
	err := db.Model(t).Update("is_delete",1).Error
	return err
}

func (t *EduTeacher) Delete() error {
	err := Remove(t)
	return err
}

func (t *EduTeacher) All() (teachers *[]EduTeacher,total uint,err error) {
	teachers = &[]EduTeacher{}
	t.IsDelete = 0
	total,err = List(t,teachers)
	return
}
