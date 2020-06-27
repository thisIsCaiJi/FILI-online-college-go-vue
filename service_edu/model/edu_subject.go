package model

import (
	"time"
)

//一级分类
type NodeSubject struct {
	Id       string        `json:"id"`
	Title    string        `json:"title"`
	Children *[]NodeSubject `json:"children"`
}

//data
type EduSubject struct {
	GmtCreate   time.Time `gorm:"column:gmt_create" form:"gmt_create" json:"gmt_create" comment:"创建时间" columnType:"datetime" dataType:"datetime" columnKey:""`
	GmtModified time.Time `gorm:"column:gmt_modified" form:"gmt_modified" json:"gmt_modified" comment:"更新时间" columnType:"datetime" dataType:"datetime" columnKey:""`
	Id          string    `gorm:"column:id" form:"id" json:"id" comment:"课程类别ID" columnType:"char(19)" dataType:"char" columnKey:"PRI"`
	ParentId    string    `gorm:"column:parent_id" form:"parent_id" json:"parent_id" comment:"父ID" columnType:"char(19)" dataType:"char" columnKey:"MUL"`
	Sort        uint      `gorm:"column:sort" form:"sort" json:"sort" comment:"排序字段" columnType:"int unsigned" dataType:"int" columnKey:""`
	Title       string    `gorm:"column:title" form:"title" json:"title" comment:"类别名称" columnType:"varchar(10)" dataType:"varchar" columnKey:""`
}

//TableName
func (s *EduSubject) TableName() string {
	return "edu_subject"
}

//query
type SubjectQuery struct {
}

//One
func (s *EduSubject) One() (one *EduSubject, err error) {
	one = &EduSubject{}
	err = One(s, one)
	return one, err
}

//List
func (s *EduSubject) ListPage(current, limit int, query SubjectQuery) (list *[]EduSubject, total uint, err error) {
	list = &[]EduSubject{}
	if err := db.Limit(limit).Offset((current - 1) * limit).Find(list).Error; err != nil {
		return nil, 0, err
	}
	if err := db.Model(s).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	return list, total, nil
}

//List
func (s *EduSubject) List() (list *[]EduSubject, total uint, err error) {
	list = &[]EduSubject{}
	List(s, list)
	return list, total, nil
}

//Update
func (s *EduSubject) Update() (err error) {
	s.GmtModified = time.Now()
	return db.Save(s).Error
}

//Create
func (s *EduSubject) Add() (err error) {
	return Create(s)
}

//Delete
func (s *EduSubject) Delete() (err error) {
	return Remove(s)
}

func (s *EduSubject) GetByTitleAndParent() (one *EduSubject, notexist bool) {
	one = &EduSubject{}
	notexist = db.Where("title = ?", s.Title).Where(" parent_id = ?", s.ParentId).First(one).RecordNotFound()
	return
}

func (s *EduSubject) GetTwoSub() (list *[]EduSubject, err error) {
	list = &[]EduSubject{}
	err = db.Where("parent_id != ?", "0").Find(list).Error
	return list, err
}
