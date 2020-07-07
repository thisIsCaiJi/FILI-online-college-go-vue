package model

import "time"

//EduCourseDescription
type EduCourseDescription struct {
	Description string    `gorm:"column:description" form:"description" json:"description" comment:"课程简介" columnType:"text" dataType:"text" columnKey:""`
	GmtCreate   time.Time `gorm:"column:gmt_create" form:"gmt_create" json:"gmt_create" comment:"创建时间" columnType:"datetime" dataType:"datetime" columnKey:""`
	GmtModified time.Time `gorm:"column:gmt_modified" form:"gmt_modified" json:"gmt_modified" comment:"更新时间" columnType:"datetime" dataType:"datetime" columnKey:""`
	Id          string    `gorm:"column:id" form:"id" json:"id" comment:"课程ID" columnType:"char(19)" dataType:"char" columnKey:"PRI"`
}

//TableName
func (m *EduCourseDescription) TableName() string {
	return "edu_course_description"
}

//One
func (m *EduCourseDescription) One() (one *EduCourseDescription, err error) {
	one = &EduCourseDescription{}
	err = One(m, one)
	return
}

//All
func (m *EduCourseDescription) All() (list *[]EduCourseDescription, total uint, err error) {
	list = &[]EduCourseDescription{}
	total, err = List(m, list)
	return
}

//page
func (m *EduCourseDescription) AllPage(current, limit int) (list *[]EduCourseDescription, total uint, err error) {
	list = &[]EduCourseDescription{}
	total, err = ListPage(m, current, limit, list)
	return
}

//Update
func (m *EduCourseDescription) Update() (err error) {
	return Update(m)
}

//CreateUserOfRole
func (m *EduCourseDescription) Create() (err error) {
	return Create(m)
}

//Delete
func (m *EduCourseDescription) Delete() (err error) {
	return Remove(m)
}
