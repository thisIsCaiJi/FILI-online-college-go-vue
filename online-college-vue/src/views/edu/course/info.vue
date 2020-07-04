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
      <el-form-item label="课程分类">
        <el-select
          v-model="courseInfo.subject_parent_id"
          clearable
          placeholder="请选择"
          @change="getTwoSubject"
        >
          <el-option
            v-for="subject in subjectOneList"
            :key="subject.id"
            :label="subject.title"
            :value="subject.id"
          ></el-option>
        </el-select>
        <el-select v-model="courseInfo.subject_id" clearable placeholder="请选择">
          <el-option
            v-for="subject in subjectTwoList"
            :key="subject.id"
            :label="subject.title"
            :value="subject.id"
          ></el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="总课时">
        <el-input-number v-model="courseInfo.lesson_num" controls-position="right" :min="0" />
      </el-form-item>
      <el-form-item label="课程简介">
        <tinymce_text v-model="courseInfo.description"></tinymce_text>
      </el-form-item>
      <el-form-item label="课程封面">
        <el-upload
          class="avatar-uploader"
          :action="baseApi+'/eduoss/fileoss'"
          :show-file-list="false"
          :on-success="handleAvatarSuccess"
          :before-upload="beforeAvatarUpload"
          accept=".jpg, .png, .gif"
        >
          <img v-if="courseInfo.cover" :src="courseInfo.cover" class="avatar" />
          <i v-else class="el-icon-plus avatar-uploader-icon"></i>
        </el-upload>
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
import subjectApi from "@/api/edu/subject";

import tinymce_text from "@/components/tinymce/index.vue";

export default {
  components: { tinymce_text },
  data() {
    return {
      saveBtnDisabled: false,
      courseInfo: {
        lesson_num: 0,
        price: 0,
        subject_id: "",
        description:"",
        cover:
          "https://fili-online-college-0001.oss-cn-shenzhen.aliyuncs.com/u%3D2496386706%2C575791735%26fm%3D26%26gp%3D0.jpg"
      },
      teacherList: [],
      subjectOneList: [],
      subjectTwoList: [],
      baseApi: process.env.VUE_APP_BASE_API
    };
  },
  created() {
    this.getListTeacher();
    this.getOneSubject();
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
    getListTeacher() {
      courseApi.getListTeacher().then(response => {
        this.teacherList = response.data.rows;
      });
    },
    getOneSubject() {
      subjectApi.getAllSubject().then(response => {
        this.subjectOneList = response.data.list;
      });
    },
    getTwoSubject(value) {
      this.subjectTwoList = [];
      this.courseInfo.subject_id = "";
      subjectApi.getAllSubject().then(response => {
        this.subjectOneList.forEach(one => {
          if (one.id === value) {
            this.subjectTwoList = one.children;
          }
        });
      });
    },
    //上传封面成功回调方法
    handleAvatarSuccess(res, file) {
      this.courseInfo.cover = res.data.url;
    },
    //上传前调用的方法
    beforeAvatarUpload(file) {
      const isImage = ["image/jpeg", "image/png", "image/gif"].includes(
        file.type
      );
      const isLt2M = file.size / 1024 / 1024 < 2;

      if (!isImage) {
        this.$message.error("上传封面图片只能是 JPG/PNG/GIF 格式!");
      }
      if (!isLt2M) {
        this.$message.error("上传头像图片大小不能超过 2MB!");
      }
      return isImage && isLt2M;
    }
  }
};
</script>