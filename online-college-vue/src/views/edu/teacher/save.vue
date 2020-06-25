<template>
  <div class="app-container">
    <el-form ref="form" label-width="120px">
      <el-form-item label="讲师名称">
        <el-input v-model="teacher.name"></el-input>
      </el-form-item>
      <el-form-item label="讲师排序">
        <el-input-number v-model="teacher.sort" controls-position="right" :min="0" />
      </el-form-item>
      <el-form-item label="讲师头衔">
        <el-select v-model="teacher.level" clearable placeholder="请选择">
          <el-option label="高级讲师" :value="1"></el-option>
          <el-option label="首席讲师" :value="2"></el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="讲师资历">
        <el-input v-model="teacher.career"></el-input>
      </el-form-item>
      <el-form-item label="讲师简介">
        <el-input v-model="teacher.intro" :rows="10" type="textarea"></el-input>
      </el-form-item>

      <el-form-item label="讲师头像">
        <el-upload
          class="avatar-uploader"
          :action="baseApi+'/eduoss/fileoss'"
          :show-file-list="false"
          :on-success="handleAvatarSuccess"
          :before-upload="beforeAvatarUpload"
        >
          <img v-if="teacher.avatar" :src="teacher.avatar" class="avatar" />
          <i v-else class="el-icon-plus avatar-uploader-icon"></i>
        </el-upload>
      </el-form-item>
      <el-form-item>
        <el-button :disabled="saveBtnDisabled" type="primary" @click="saveOrUpdate()">保存</el-button>
        <router-link :to="'/teacher/table'">
          <el-button>取消</el-button>
        </router-link>
      </el-form-item>
    </el-form>
  </div>
</template>

<style>
.avatar-uploader .el-upload {
  border: 1px dashed #d9d9d9;
  border-radius: 6px;
  cursor: pointer;
  position: relative;
  overflow: hidden;
}
.avatar-uploader .el-upload:hover {
  border-color: #409eff;
}
.avatar-uploader-icon {
  font-size: 28px;
  color: #8c939d;
  width: 178px;
  height: 178px;
  line-height: 178px;
  text-align: center;
}
.avatar {
  width: 178px;
  height: 178px;
  display: block;
}
</style>

<script>
import teacherApi from "@/api/edu/teacher";

export default {
  data() {
    return {
      baseApi: process.env.VUE_APP_BASE_API,
      teacher: { sort: 0, level: 1, avatar: null },
      saveBtnDisabled: false,
    };
  },
  created() {
    this.init();
  },
  watch: {
    $route(to, from) {
      this.init();
    }
  },
  methods: {
    saveOrUpdate() {
      this.saveBtnDisabled = true;
      if (this.teacher.id) {
        this.updateTeacher();
      } else {
        this.saveTeacher();
      }
      this.saveBtnDisabled = false;
    },
    saveTeacher() {
      teacherApi
        .addTeacher(this.teacher)
        .then(response => {
          this.$message({
            message: "添加讲师成功",
            type: "success"
          });
          this.$router.push({ path: "/teacher/table" });
        })
        .catch(error => {
          this.$message.error("添加讲师失败");
        });
    },
    getTeacherById(id) {
      teacherApi.getTeacherById(id).then(response => {
        this.teacher = response.data.teacher;
      });
    },
    updateTeacher() {
      teacherApi
        .updateTeacher(this.teacher)
        .then(response => {
          this.$message({
            message: "修改讲师成功",
            type: "success"
          });
          this.$router.push({ path: "/teacher/table" });
        })
        .catch(error => {
          this.$message.error("修改讲师失败");
        });
    },
    init() {
      if (this.$route.params && this.$route.params.id) {
        const id = this.$route.params.id;
        this.getTeacherById(id);
      } else {
        this.teacher = { sort: 0, level: 1, avatar: null  };
      }
    },
    handleAvatarSuccess(res, file) {
      this.teacher.avatar = res.data.url;
    },
    beforeAvatarUpload(file) {
      const isImage = ["image/jpeg","image/png","image/gif"].includes(file.type)
      const isLt2M = file.size / 1024 / 1024 < 2;

      if (!isImage) {
        this.$message.error("上传头像图片只能是 JPG 格式!");
      }
      if (!isLt2M) {
        this.$message.error("上传头像图片大小不能超过 2MB!");
      }
      return isImage && isLt2M;
    }
  } 
};
</script>