<template>
  <div class="app-container">
    <el-steps :active="2" finish-status="wait">
      <el-step title="填写课程基本信息"></el-step>
      <el-step title="创建课程大纲"></el-step>
      <el-step title="最终发布"></el-step>
    </el-steps>

    <el-button style="margin-top: 12px;" @click="previous">上一步</el-button>
    <el-button type="primary" style="margin-top: 12px;" @click="next">发布</el-button>
  </div>
</template>

<script>
// import courseApi  from "@/api/edu/course";

export default {
  data() {
    return {
        
    };
  },
  created() {},
  methods: {
      previous(){
          this.$router.push({path:'/course/chapter/1'})
      },
      next(){
          this.$router.push({path:'/course/list'})
      }
  }
};
</script>