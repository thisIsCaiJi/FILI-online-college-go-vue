package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(Cors())
	user := router.Group("/api/eduservice/user")
	{
		user.POST("/login", Login)
		user.GET("/info", Info)
		user.POST("/logout", Logout)
	}
	video := router.Group("/api/eduservice/video")
	{
		video.POST("/addVideo", AddVideo)
		video.GET("/:id", GetVideo)
		video.DELETE("/:id", RemoveVideo)
		video.POST("/updateVideo", UpdateVideo)
	}
	teacher := router.Group("/api/eduservice/edu-teacher")
	{
		teacher.POST("/pageTeacherCondition/:current/:limit", PageTeacherConditionPost)
		teacher.POST("/addTeacher", AddTeacher)
		teacher.GET("/getTeacher/:id", GetTeacherById)
		teacher.POST("/updateTeacher", UpdateTeacher)
		teacher.DELETE("/:id", RemoveTeacher)
		teacher.GET("/all", TeacherList)
	}
	subject := router.Group("api/eduservice/subject")
	{
		subject.POST("/import",ImportSubject)
		subject.GET("/allSubject",GetAllSubject)
	}
	course := router.Group("api/eduservice/course")
	{
		course.POST("/addCourseInfo",AddCourseInfo)
		course.GET("/getCourseInfo/:id",GetCourseInfo)
		course.POST("/updateCourseInfo",UpdateCourseInfo)
		course.GET("/coursePublish/:id",GetCoursePublishVo)
	}
	chapter := router.Group("api/eduservice/chapter")
	{
		chapter.GET("/chapterVideo/:course_id",getChapterVideo)
		chapter.POST("/addChapter",AddChapter)
		chapter.GET("/getChapter/:id",GetChapterInfo)
		chapter.POST("/updateChapter",UpdateChapterInfo)
		chapter.DELETE("/:id",DeleteChapterInfo)
	}
	return router
}

// 处理跨域请求,支持options访问
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		fmt.Printf("origin:%s\n", c.GetHeader("Access-Control-Allow-Origin"))
		// 放行所有OPTIONS方法，因为有的模板是要请求两次的
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}
