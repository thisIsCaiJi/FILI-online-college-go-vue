package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/thisIsCaiJi/online_college/service_edu/model"
)

func init(){
	group.GET("/eduservice/chapterVideo/:course_id",getChapterVideo)
	group.POST("/eduservice/chapter/addChapter",AddChapter)
	group.GET("/eduservice/chapter/:id",GetChapterInfo)
	group.POST("/eduservice/chapter/updateChapter",UpdateChapterInfo)
	group.DELETE("/eduservice/chapter/:id",DeleteChapterInfo)
}

func getChapterVideo(ctx *gin.Context){

	courseId := ctx.Param("course_id")
	if courseId == "" {
		HandlerError(nil,"获取参数失败")
		JsonError(ctx)
		return
	}
	chapter := &model.EduChapter{CourseId:courseId}
	video := &model.EduVideo{CourseId:courseId}
	chapterList,_,err := chapter.All()
	if err != nil {
		HandlerError(err,"查询chapter失败")
		JsonError(ctx)
		return
	}
	videoList,_,err := video.All()
	if err != nil {
		HandlerError(err,"查询video失败")
		JsonError(ctx)
		return
	}
	list := make([]*model.ChapterVo,0,len(*chapterList))
	m := make(map[string]*[]model.VideoVo)
	for _,chapt := range *chapterList{
		one := &model.ChapterVo{
			Id:chapt.Id,
			Title:chapt.Title,
			Children:&[]model.VideoVo{},
		}
		list = append(list,one)
		m[chapt.Id] = one.Children
	}
	for _,vid := range *videoList{
		if _,ok := m[vid.ChapterId]; ok{
			one := &model.VideoVo{Id: vid.Id, Title: vid.Title}
			*m[vid.ChapterId] = append(*m[vid.ChapterId],*one)
		}
	}
	JsonSuccessKV(ctx,"allChapterVideo",list)
}

func AddChapter(ctx *gin.Context){
	chapter := &model.EduChapter{}
	ctx.ShouldBindJSON(chapter)
	if err := chapter.Create();err!=nil {
		JsonError(ctx)
	}
	JsonSuccess(ctx)
}

func GetChapterInfo(ctx *gin.Context){
	id := ctx.Param("id")
	chapterParams := &model.EduChapter{Id:id}
	chapter,err := chapterParams.One()
	if err != nil {
		JsonError(ctx)
		return
	}
	JsonSuccessKV(ctx,"chapter",chapter)
}

func UpdateChapterInfo(ctx *gin.Context){
	chapter := &model.EduChapter{}
	ctx.ShouldBindJSON(chapter)
	if err := chapter.Update();err!=nil {
		JsonError(ctx)
		return
	}
	JsonSuccess(ctx)
}

func DeleteChapterInfo(ctx *gin.Context){
	id := ctx.Param("id")
	chapter := &model.EduChapter{Id:id}
	video := &model.EduVideo{ChapterId:id}
	if total := video.Count();total > 0 {
		JsonErrorMessage(ctx,"该章节下存在小节")
		return
	}
	if err := chapter.Delete();err!=nil {
		JsonErrorMessage(ctx,"删除章节失败")
		return
	}
	JsonSuccess(ctx)
}