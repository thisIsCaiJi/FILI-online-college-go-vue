import request from '@/utils/request'

export default {
    getChapterVideo(courseInfo) {
        return request({
            url: `/eduservice/chapter/chapterVideo/${courseInfo}`,
            method: 'get',
        })
    },
    addChapter(chapter) {
        return request({
            url: `/eduservice/chapter/addChapter`,
            method: 'post',
            data: chapter
        })
    },
    getChapter(chapterId){
        return request({
            url: `/eduservice/chapter/getChapter/${chapterId}`,
            method: 'get',
        })
    },
    updateChapter(chapter){
        return request({
            url: `/eduservice/chapter/updateChapter`,
            method: 'post',
            data: chapter
        })
    },
    deleteChapter(chapterId){
        return request({
            url: `/eduservice/chapter/${chapterId}`,
            method: 'delete',
        })
    },
}
