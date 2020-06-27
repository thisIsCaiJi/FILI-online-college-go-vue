package util

import (
	"fmt"
	"testing"
	"time"
)

type EduCourseVo struct {
	BuyCount        uint64    `json:"buy_count"`
	Cover           string    `json:"cover"`
	LessonNum       uint      `json:"lesson_num"`
	Price           float32   `json:"price"`
	Status          string    `json:"status"`
	SubjectId       string    `json:"subject_id"`
	SubjectParentId string    `json:"subject_parent_id"`
	TeacherId       string    `json:"teacher_id"`
	Title           string    `json:"title"`
	Version         uint64    `json:"version"`
	ViewCount       uint64    `json:"view_count"`
	Description 	string	  `json:"description"`
}

type EduCourse struct {
	BuyCount        uint64    `gorm:"column:buy_count" form:"buy_count" json:"buy_count" comment:"销售数量" columnType:"bigint unsigned" dataType:"bigint" columnKey:""`
	Cover           string    `gorm:"column:cover" form:"cover" json:"cover" comment:"课程封面图片路径" columnType:"varchar(255)" dataType:"varchar" columnKey:""`
	GmtCreate       time.Time `gorm:"column:gmt_create" form:"gmt_create" json:"gmt_create" comment:"创建时间" columnType:"datetime" dataType:"datetime" columnKey:""`
	GmtModified     time.Time `gorm:"column:gmt_modified" form:"gmt_modified" json:"gmt_modified" comment:"更新时间" columnType:"datetime" dataType:"datetime" columnKey:""`
	Id              string    `gorm:"column:id" form:"id" json:"id" comment:"课程ID" columnType:"char(19)" dataType:"char" columnKey:"PRI"`
	IsDeleted       int       `gorm:"column:is_deleted" form:"is_deleted" json:"is_deleted" comment:"逻辑删除 1（true）已删除， 0（false）未删除" columnType:"tinyint" dataType:"tinyint" columnKey:""`
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

func Test_CopyStruct(t *testing.T) {
	a := EduCourse{
		BuyCount:        0,
		Cover:           "",
		GmtCreate:       time.Time{},
		GmtModified:     time.Time{},
		Id:              "",
		IsDeleted:       0,
		LessonNum:       0,
		Price:           0,
		Status:          "",
		SubjectId:       "",
		SubjectParentId: "",
		TeacherId:       "",
		Title:           "",
		Version:         0,
		ViewCount:       0,
	}
	b := EduCourseVo{
		BuyCount:        1,
		Cover:           "Status",
		LessonNum:       2,
		Price:           1,
		Status:          "Status",
		SubjectId:       "SubjectId",
		SubjectParentId: "SubjectParentId",
		TeacherId:       "TeacherId",
		Title:           "23132",
		Version:        2,
		ViewCount:       1,
		Description:     "12321321",
	}
	CopyStruct(&b,&a)
	fmt.Printf("a:%v,b%v\n",a,b)
}