import request from '@/utils/request'

export default {
    addVideo(video) {
        return request({
            url: `/eduservice/video/addVideo`,
            method: 'post',
            data: video
        })
    },
    deleteVideo(id) {
        return request({
            url: `/eduservice/video/${id}`,
            method: 'delete',
        })
    },
    getVideo(id) {
        return request({
            url: `/eduservice/video/${id}`,
            method: 'get',
        })
    },
    updateVideo(video) {
        return request({
            url: `/eduservice/video/updateVideo`,
            method: 'post',
            data: video
        })
    },
    deleteAliyunvod(id) {
        return request({
            url: `/eduvod/video/${id}`,
            method: 'delete',
        })
    }
}