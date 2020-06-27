package model

import "time"

//EduVideo
type EduVideo struct {
	ChapterId         string    `gorm:"column:chapter_id" form:"chapter_id" json:"chapter_id" comment:"章节ID" columnType:"char(19)" dataType:"char" columnKey:"MUL"`
	CourseId          string    `gorm:"column:course_id" form:"course_id" json:"course_id" comment:"课程ID" columnType:"char(19)" dataType:"char" columnKey:"MUL"`
	Duration          float32   `gorm:"column:duration" form:"duration" json:"duration" comment:"视频时长（秒）" columnType:"float" dataType:"float" columnKey:""`
	GmtCreate         time.Time `gorm:"column:gmt_create" form:"gmt_create" json:"gmt_create" comment:"创建时间" columnType:"datetime" dataType:"datetime" columnKey:""`
	GmtModified       time.Time `gorm:"column:gmt_modified" form:"gmt_modified" json:"gmt_modified" comment:"更新时间" columnType:"datetime" dataType:"datetime" columnKey:""`
	Id                string    `gorm:"column:id" form:"id" json:"id" comment:"视频ID" columnType:"char(19)" dataType:"char" columnKey:"PRI"`
	IsFree            uint      `gorm:"column:is_free" form:"is_free" json:"is_free" comment:"是否可以试听：0收费 1免费" columnType:"tinyint unsigned" dataType:"tinyint" columnKey:""`
	PlayCount         uint64    `gorm:"column:play_count" form:"play_count" json:"play_count" comment:"播放次数" columnType:"bigint unsigned" dataType:"bigint" columnKey:""`
	Size              uint64    `gorm:"column:size" form:"size" json:"size" comment:"视频源文件大小（字节）" columnType:"bigint unsigned" dataType:"bigint" columnKey:""`
	Sort              uint      `gorm:"column:sort" form:"sort" json:"sort" comment:"排序字段" columnType:"int unsigned" dataType:"int" columnKey:""`
	Status            string    `gorm:"column:status" form:"status" json:"status" comment:"Empty未上传 Transcoding转码中  Normal正常" columnType:"varchar(20)" dataType:"varchar" columnKey:""`
	Title             string    `gorm:"column:title" form:"title" json:"title" comment:"节点名称" columnType:"varchar(50)" dataType:"varchar" columnKey:""`
	Version           uint64    `gorm:"column:version" form:"version" json:"version" comment:"乐观锁" columnType:"bigint unsigned" dataType:"bigint" columnKey:""`
	VideoOriginalName string    `gorm:"column:video_original_name" form:"video_original_name" json:"video_original_name" comment:"原始文件名称" columnType:"varchar(100)" dataType:"varchar" columnKey:""`
	VideoSourceId     string    `gorm:"column:video_source_id" form:"video_source_id" json:"video_source_id" comment:"云端视频资源" columnType:"varchar(100)" dataType:"varchar" columnKey:""`
}

//TableName
func (m *EduVideo) TableName() string {
	return "edu_video"
}

//One
func (m *EduVideo) One() (one *EduVideo, err error) {
	one = &EduVideo{}
	err = One(m, one)
	return
}

//All
func (m *EduVideo) All() (list *[]EduVideo, total uint, err error) {
	list = &[]EduVideo{}
	total, err = List(m, list)
	return
}

//page
func (m *EduVideo) AllPage() (list *[]EduVideo, current, limit int, total uint, err error) {
	list = &[]EduVideo{}
	total, err = ListPage(m, current, limit, list)
	return
}

//Update
func (m *EduVideo) Update() (err error) {
	return Update(m)
}

//CreateUserOfRole
func (m *EduVideo) Create() (err error) {
	return Create(m)
}

//Delete
func (m *EduVideo) Delete() (err error) {
	return Remove(m)
}
