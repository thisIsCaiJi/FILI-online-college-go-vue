package model

import "time"


//EduChapter
type EduChapter struct {
	CourseId    string    `gorm:"column:course_id" form:"course_id" json:"course_id" comment:"课程ID" columnType:"char(19)" dataType:"char" columnKey:"MUL"`
	GmtCreate   time.Time `gorm:"column:gmt_create" form:"gmt_create" json:"gmt_create" comment:"创建时间" columnType:"datetime" dataType:"datetime" columnKey:""`
	GmtModified time.Time `gorm:"column:gmt_modified" form:"gmt_modified" json:"gmt_modified" comment:"更新时间" columnType:"datetime" dataType:"datetime" columnKey:""`
	Id          string    `gorm:"column:id" form:"id" json:"id" comment:"章节ID" columnType:"char(19)" dataType:"char" columnKey:"PRI"`
	Sort        uint      `gorm:"column:sort" form:"sort" json:"sort" comment:"显示排序" columnType:"int unsigned" dataType:"int" columnKey:""`
	Title       string    `gorm:"column:title" form:"title" json:"title" comment:"章节名称" columnType:"varchar(50)" dataType:"varchar" columnKey:""`
}

//TableName
func (m *EduChapter) TableName() string {
	return "edu_chapter"
}

//One
func (m *EduChapter) One() (one *EduChapter, err error) {
	one = &EduChapter{}
	err = One(m, one)
	return
}

//All
func (m *EduChapter) All() (list *[]EduChapter, total uint, err error) {
	list = &[]EduChapter{}
	total, err = List(m, list)
	return
}

//page
func (m *EduChapter) AllPage() (list *[]EduChapter, current, limit int, total uint, err error) {
	list = &[]EduChapter{}
	total, err = ListPage(m, current, limit, list)
	return
}

//Update
func (m *EduChapter) Update() (err error) {
	return Update(m)
}

//CreateUserOfRole
func (m *EduChapter) Create() (err error) {
	return Create(m)
}

//Delete
func (m *EduChapter) Delete() (err error) {
	return Remove(m)
}


type ChapterVo struct {
	Id string`json:"id"`
	Title string`json:"title"`
	Children *[]VideoVo`json:"children"`
}

func OneChapterVo() *ChapterVo {
	return &ChapterVo{Children:&[]VideoVo{}}
}