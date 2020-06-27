import request from '@/utils/request'

export default{
    getAllSubject() {
        return request({
          url: `/eduservice/subject/allSubject`,
          method: 'get',
        })
    },
}
