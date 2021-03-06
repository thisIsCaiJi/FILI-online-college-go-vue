import request from '@/utils/request'

export default{
    getTeacherListPage(current,limit,teacherQuery) {
        return request({
          url: `/eduservice/edu-teacher/pageTeacherCondition/${current}/${limit}`,
          method: 'post',
          data: teacherQuery
        })
    },
    removeTeacherById(id) {
      return request({
        url: `/eduservice/edu-teacher/${id}`,
        method: 'delete'
      })
    },
    addTeacher(teacher) {
      return request({
        url: `/eduservice/edu-teacher/addTeacher`,
        method: 'post',
        data: teacher
      })
    },
    getTeacherById(id) {
      return request({
        url: `/eduservice/edu-teacher/getTeacher/${id}`,
        method: 'get'
      })
    },
    updateTeacher(teacher) {
      return request({
        url: `/eduservice/edu-teacher/updateTeacher`,
        method: 'post',
        data: teacher
      })
    }
}
