package model

import (
	"github.com/thisIsCaiJi/online_college/service_edu/util"
	"time"
)

type EduCourseVo struct {
	Id              string  `json:"id"`
	BuyCount        uint64  `json:"buy_count"`
	Cover           string  `json:"cover"`
	LessonNum       uint    `json:"lesson_num"`
	Price           float32 `json:"price"`
	Status          string  `json:"status"`
	SubjectId       string  `json:"subject_id"`
	SubjectParentId string  `json:"subject_parent_id"`
	TeacherId       string  `json:"teacher_id"`
	Title           string  `json:"title"`
	Version         uint64  `json:"version"`
	ViewCount       uint64  `json:"view_count"`
	Description     string  `json:"description"`
}

type EduCourseQuery struct {
	Status          string  `json:"status"`
	Title           string  `json:"title"`
}

const CourseStatusDraft = "Draft"
const CourseStatusNormal = "Normal"

// 和模型关联的结构体
type EduCourse struct {
	BuyCount        uint64    `gorm:"column:buy_count" form:"buy_count" json:"buy_count" comment:"销售数量" columnType:"bigint unsigned" dataType:"bigint" columnKey:""`
	Cover           string    `gorm:"column:cover" form:"cover" json:"cover" comment:"课程封面图片路径" columnType:"varchar(255)" dataType:"varchar" columnKey:""`
	GmtCreate       time.Time `gorm:"column:gmt_create" form:"gmt_create" json:"gmt_create" comment:"创建时间" columnType:"datetime" dataType:"datetime" columnKey:""`
	GmtModified     time.Time `gorm:"column:gmt_modified" form:"gmt_modified" json:"gmt_modified" comment:"更新时间" columnType:"datetime" dataType:"datetime" columnKey:""`
	Id              string    `gorm:"column:id" form:"id" json:"id" comment:"课程ID" columnType:"char(19)" dataType:"char" columnKey:"PRI"`
	LessonNum       uint      `gorm:"column:lesson_num" form:"lesson_num" json:"lesson_num" comment:"总课时" columnType:"int unsigned" dataType:"int" columnKey:""`
	Price           float32   `gorm:"column:price" form:"price" json:"price" comment:"课程销售价格，设置为0则可免费观看" columnType:"decimal(10,2) unsigned" dataType:"decimal" columnKey:""`
	Status          string    `gorm:"column:status" form:"status" json:"status" comment:"课程状态 Draft未发布  Normal已发布" columnType:"varchar(10)" dataType:"varchar" columnKey:""`
	SubjectId       string    `gorm:"column:subject_id" form:"subject_id" json:"subject_id" comment:"课程专业ID" columnType:"char(19)" dataType:"char" columnKey:"MUL"`
	SubjectParentId string    `gorm:"column:subject_parent_id" form:"subject_parent_id" json:"subject_parent_id" comment:"课程专业父级ID" columnType:"char(19)" dataType:"char" columnKey:""`
	TeacherId       string    `gorm:"column:teacher_id" form:"teacher_id" json:"teacher_id" comment:"课程讲师ID" columnType:"char(19)" dataType:"char" columnKey:"MUL"`
	Title           string    `gorm:"column:title" form:"title" json:"title" comment:"课程标题" columnType:"varchar(50)" dataType:"varchar" columnKey:"MUL"`
	Version         uint64    `gorm:"column:version" form:"version" json:"version" comment:"乐观锁" columnType:"bigint unsigned" dataType:"bigint" columnKey:""`
	ViewCount       uint64    `gorm:"column:view_count" form:"view_count" json:"view_count" comment:"浏览数量" columnType:"bigint unsigned" dataType:"bigint" columnKey:""`
}

//TableName
func (m *EduCourse) TableName() string {
	return "edu_course"
}

//One
func (m *EduCourse) One() (one *EduCourse, err error) {
	one = &EduCourse{}
	err = One(m, one)
	return
}

//All
func (m *EduCourse) All() (list *[]EduCourse, total uint, err error) {
	list = &[]EduCourse{}
	total, err = List(m, list)
	return
}

//page
func (m *EduCourse) AllPage(current, limit int) (list *[]EduCourse, total uint, err error) {
	list = &[]EduCourse{}
	total, err = ListPage(m, current, limit, list)
	return
}

//page
func (m *EduCourse) ListPage(current, limit int, query *EduCourseQuery) (list *[]EduCourse, total uint, err error) {
	list = &[]EduCourse{}
	theDb := db.Model(m)
	if query.Status == CourseStatusNormal {
		theDb = theDb.Where("status = ?",CourseStatusNormal)
	}else if query.Status != "" {
		theDb = theDb.Where("status != ?",CourseStatusNormal)
	}
	if query.Title != "" {
		theDb = theDb.Where("title like ?","%"+query.Title+"%")
	}
	err = theDb.Count(&total).Limit(limit).Offset((current-1) * limit).Find(list).Error
	return
}

//Update
func (m *EduCourse) Update() (err error) {
	return Update(m)
}

//Create
func (m *EduCourse) Create() (err error) {
	return Create(m)
}

//Delete
func (m *EduCourse) Delete() (err error) {
	return Remove(m)
}

//同时创建简介
func (m *EduCourse) CreateAll(description *EduCourseDescription) (err error) {
	tx := db.Begin()
	if err = Create(m, tx); err != nil {
		tx.Rollback()
		return
	}
	description.Id = m.Id
	if err = Create(description, tx); err != nil {
		tx.Rollback()
		return
	}
	return tx.Commit().Error
}

// 同时查出简介
func (m *EduCourse) OneAll() (one *EduCourseVo, err error) {
	course, courseDescription, description := &EduCourse{}, &EduCourseDescription{Id: m.Id}, &EduCourseDescription{}
	if err = One(m, course); err != nil {
		return nil, err
	}
	courseVo := &EduCourseVo{}
	util.CopyStruct(course, courseVo)
	if err = One(courseDescription, description); err == nil {
		courseVo.Description = description.Description
	}
	return courseVo, nil
}

//同时修改简介
func (m *EduCourse) UpdateAll(description *EduCourseDescription) (err error) {
	tx := db.Begin()
	if err = Update(m, tx); err != nil {
		tx.Rollback()
		return
	}
	description.Id = m.Id
	if err = Update(description, tx); err != nil {
		tx.Rollback()
		return
	}
	return tx.Commit().Error
}

// 删除课程和课程的章节、课时、描述
func (m *EduCourse) RemoveAll() (err error) {
	tx := db.Begin()
	// 删除课时
	if err = Remove(&EduVideo{CourseId:m.Id},tx); err != nil {
		tx.Rollback()
		return err
	}
	// 删除章节
	if err = Remove(&EduChapter{CourseId:m.Id},tx); err != nil {
		tx.Rollback()
		return err
	}
	// 删除描述
	if err = Remove(&EduCourseDescription{Id:m.Id},tx); err != nil {
		tx.Rollback()
		return err
	}
	// 删除课程
	if err = Remove(m,tx); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

type CoursePublishVo struct {
	Title       string `json:"title"`
	Cover       string `json:"cover"`
	LessonNum   int    `json:"lesson_num"`
	OneSubject  string `json:"one_subject"`
	TwoSubject  string `json:"two_subject"`
	TeacherName string `json:"teacher_name"`
	Price       string `json:"price"`
	Description string `json:"description"`
}

func (m *EduCourse) GetCoursePublishVo() (vo *CoursePublishVo, err error) {
	vo = &CoursePublishVo{}
	err = db.Raw("SELECT ec.id, ec.title, ec.price, ec.lesson_num, ec.cover, ecd.description, et.`name` AS teacher_name, es1.title AS two_subject, es2.title AS one_subject FROM edu_course ec LEFT JOIN edu_course_description ecd ON ec.id = ecd.id LEFT JOIN edu_teacher et ON ec.teacher_id = et.id LEFT JOIN edu_subject es1 ON ec.subject_id = es1.id LEFT JOIN edu_subject es2 ON ec.subject_parent_id = es2.id WHERE ec.id = ? limit 1", m.Id).Scan(vo).Error
	return
}
