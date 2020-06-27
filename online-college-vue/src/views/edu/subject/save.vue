<template>
  <div class="app-container">
    <el-from label-width="120px">
      <el-from-item label="信息描述">
        <el-tag type="info">excel模板说明</el-tag>
        <el-tag>
          <i class="el-icon-download" />
          <a :href="'../static/01.xlsx'">点击下载模板</a>
        </el-tag>
      </el-from-item>
      <el-from-item label="选择Excel">
        <el-upload
          class="upload-demo"
          ref="upload"
          :action="baseApi+'/eduservice/subject/import'"
          :on-success="uploadSuccess"
          :on-error="uploadError"
          :file-list="fileList"
          :auto-upload="false"
          :limit="1"
          name="file"
          accept=".xlsx"
          :disabled="isUploading"
        >
          <el-button slot="trigger" size="small" type="primary" :loading="isUploading">选取文件</el-button>
          <el-button
            style="margin-left: 10px;"
            size="small"
            type="success"
            @click="submitUpload"
            :loading="isUploading"
          >上传到服务器</el-button>
          <div slot="tip" class="el-upload__tip">只能上传xlsx文件，且不超过5mb</div>
        </el-upload>
      </el-from-item>
    </el-from>
  </div>
</template>


<script>
import subjectApi from "@/api/edu/subject";

export default {
  data() {
    return {
      baseApi: process.env.VUE_APP_BASE_API,
      teacher: { sort: 0, level: 1, avatar: null },
      isUploading: false
    };
  },
  watch: {
    $route(to, from) {
      this.init();
    }
  },
  methods: {
    submitUpload(){
        this.isUploading = true
        this.$refs.upload.submit()
    },
    //上传成功
    uploadSuccess() {
        this.isUploading = false
        this.$message({
            type:'success',
            message:'添加课程分类成功'
        })
        //跳转到课程分类页面
        this.$router.push({path:'/subject/list'})
    },
    //上传失败
    uploadError(){
        this.isUploading = false
        this.$message({
            type:'error',
            message:'添加课程分类失败'
        })
    },
  }
};
</script>