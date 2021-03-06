import request from '@/utils/request'

export default {
  addCourseInfo(courseInfo) {
    return request({
      url: `/eduservice/course/addCourseInfo`,
      method: 'post',
      data: courseInfo
    })
  },
  getListTeacher() {
    return request({
      url: `/eduservice/edu-teacher/all`,
      method: 'get',
    })
  },
  getCourseInfo(id) {
    return request({
      url: `/eduservice/course/getCourseInfo/${id}`,
      method: 'get',
    })
  },
  updateCourseInfo(courseInfo) {
    return request({
      url: `/eduservice/course/updateCourseInfo`,
      method: 'post',
      data: courseInfo
    })
  },
  getCoursePublish(id) {
    return request({
      url: `/eduservice/course/coursePublish/${id}`,
      method: 'get',
    })
  },
  publishCourse(id) {
    return request({
      url: `/eduservice/course/publish/${id}`,
      method: 'put',
    })
  },
  getCourseList(current, limit, courseQuery) {
    return request({
      url: `/eduservice/course/list/${current}/${limit}`,
      method: 'post',
      data: courseQuery
    })
  },
  removeCourseInfoById(id) {
    return request({
      url: `/eduservice/course/courseInfo/${id}`,
      method: 'delete',
    })
  }
}
