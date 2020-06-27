<template>
  <div class="app-container">
    <el-steps :active="0" finish-status="wait">
      <el-step title="填写课程基本信息"></el-step>
      <el-step title="创建课程大纲"></el-step>
      <el-step title="最终发布"></el-step>
    </el-steps>

    <el-form ref="form" label-width="120px">
      <el-form-item label="课程标题">
        <el-input v-model="courseInfo.title"></el-input>
      </el-form-item>
      <el-form-item label="课程讲师">
        <el-select v-model="courseInfo.teacher_id" clearable placeholder="请选择">
          <el-option
            v-for="teacher in teacherList"
            :key="teacher.id"
            :label="teacher.name"
            :value="teacher.id"
          ></el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="总课时">
        <el-input-number v-model="courseInfo.lesson_num" controls-position="right" :min="0" />
      </el-form-item>
      <el-form-item label="课程简介">
        <el-input v-model="courseInfo.description"></el-input>
      </el-form-item>
      <el-form-item label="课程价格">
        <el-input-number v-model="courseInfo.price" controls-position="right" :min="0" />元
      </el-form-item>
    </el-form>

    <el-button
      :disabled="saveBtnDisabled"
      type="primary"
      style="margin-top: 12px;"
      @click="saveAndNext"
    >保存并下一步</el-button>
  </div>
</template>

<script>
import courseApi from "@/api/edu/course";

export default {
  data() {
    return {
      saveBtnDisabled: false,
      courseInfo: { lesson_num: 0, price: 0 },
      teacherList:[]
    };
  },
  created() {
      this.getListTeacher()
  },
  methods: {
    saveAndNext() {
      courseApi.addCourseInfo(this.courseInfo).then(response => {
        this.$message({
          message: "添加课程信息成功",
          type: "success"
        });
        var id = response.data.id;
        this.$router.push({ path: `/course/chapter/${id}` });
      });
    },
    getListTeacher(){
        courseApi.getListTeacher().then(response => {
            this.teacherList = response.data.rows
        })
    }
  }
};
</script>