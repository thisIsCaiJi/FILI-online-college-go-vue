<template>
  <div class="app-container">
    <el-steps :active="2" finish-status="wait">
      <el-step title="填写课程基本信息"></el-step>
      <el-step title="创建课程大纲"></el-step>
      <el-step title="最终发布"></el-step>
    </el-steps>
    <div class="ccInfo">
      <img :src="coursePublish.cover" />
      <div class="main">
        <h2>{{coursePublish.title}}</h2>
      </div>
    </div>
    <el-button style="margin-top: 12px;" @click="previous">上一步</el-button>
    <el-button type="primary" style="margin-top: 12px;" @click="next">发布</el-button>
  </div>
</template>

<script>
import courseApi  from "@/api/edu/course";

export default {
  data() {
    return {
      courseId: "",
      coursePublish: {title:''},
    };
  },
  created() {
    if (this.$route.params && this.$route.params.id) {
      this.courseId = this.$route.params.id;
      this.getCoursePublishById(this.courseId)
    }
  },
  methods: {
    getCoursePublishById(id) {
      courseApi.getCoursePublish(id).then(response => {
        this.coursePublish = response.data.coursePublish
      })
    },
    previous() {
      this.$router.push({ path: `/course/chapter/${this.courseId}` });
    },
    next() {
      this.$router.push({ path: "/course/list" });
    }
  }
};
</script>