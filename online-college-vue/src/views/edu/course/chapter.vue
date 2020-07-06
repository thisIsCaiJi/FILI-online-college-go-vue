<template>
  <div class="app-container">
    <el-steps :active="1" finish-status="wait">
      <el-step title="填写课程基本信息"></el-step>
      <el-step title="创建课程大纲"></el-step>
      <el-step title="最终发布"></el-step>
    </el-steps>

    <el-button type="text" @click="openChapterDialog()">添加章节</el-button>

    <ul class="chanpterList">
      <li v-for="chapter in chapterVideoList" :key="chapter.id">
        <p>
          {{chapter.title}}
          <span class="acts">
            <el-button type="text" @click="openVideoDialog(chapter.id)">添加小节</el-button>
            <el-button style type="text" @click="openEditChapterDialog(chapter.id)">编辑</el-button>
            <el-button type="text" @click="removeChapter(chapter.id)">删除</el-button>
          </span>
        </p>
        <ul class="chanpterList videoList">
          <li v-for="video in chapter.children" :key="video.id">
            <p>
              {{video.title}}
              <span class="acts">
                <el-button style type="text" @click="openEditVideoDialog(video.id)">编辑</el-button>
                <el-button type="text" @click="removeVideo(video.id)">删除</el-button>
              </span>
            </p>
          </li>
        </ul>
      </li>
    </ul>

    <el-button style="margin-top: 12px; " @click="previous">上一步</el-button>
    <el-button
      :disabled="saveBtnDisabled"
      type="primary"
      style="margin-top: 12px;"
      @click="next"
    >保存并下一步</el-button>

    <el-dialog title="添加章节" :visible.sync="dialogCHapterFormVisible">
      <el-form :model="chapter" label-width="120px">
        <el-form-item label="章节标题">
          <el-input v-model="chapter.title" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="章节排序">
          <el-input-number v-model="chapter.sort" :min="0" controls-position="right" />
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogCHapterFormVisible = false">取 消</el-button>
        <el-button type="primary" @click="saveOrUpdate()">确 定</el-button>
      </div>
    </el-dialog>

    <el-dialog title="添加课时" :visible.sync="dialogVideoFormVisible">
      <el-form :model="chapter" label-width="120px">
        <el-form-item label="课时标题">
          <el-input v-model="video.title" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="课时排序">
          <el-input-number v-model="video.sort" :min="0" controls-position="right" />
        </el-form-item>
        <el-form-item label="是否免费">
          <el-radio-group v-model="video.is_free">
            <el-radio :label="1">是</el-radio>
            <el-radio :label="0">否</el-radio>
          </el-radio-group>
          <el-form-item label="上传视频">
            <!-- TODO -->
          </el-form-item>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogVideoFormVisible = false">取 消</el-button>
        <el-button type="primary" @click="saveOrUpdateVideo()">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import chapterApi from "@/api/edu/chapter";
import videoApi from "@/api/edu/video";

export default {
  data() {
    return {
      dialogVideoFormVisible: false,
      dialogCHapterFormVisible: false,
      saveBtnDisabled: false,
      chapterVideoList: [],
      courseId: "", //课程id
      chapter: { title: "", sort: 0 },
      video: { title: "", sort: 0 }
    };
  },
  created() {
    if (this.$route.params && this.$route.params.id) {
      this.courseId = this.$route.params.id;
      this.chapter.course_id = this.$route.params.id;
      this.video.course_id = this.courseId;
      this.getChapterVideo();
    }
  },
  methods: {
    // 打开编辑课时窗口
    openEditVideoDialog(id) {
      videoApi.getVideo(id).then(response => {
        this.video = response.data.video;
        this.dialogVideoFormVisible = true;
      });
    },
    // 删除课时
    removeVideo(id) {
      this.$confirm("此操作将删除课时, 是否继续?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning"
      })
        .then(() => {
          videoApi.deleteVideo(id).then(response => {
            this.$message({
              type: "success",
              message: "删除课时成功!"
            });
            this.getChapterVideo();
          });
        })
        .catch(() => {
          this.$message({
            type: "info",
            message: "已取消删除"
          });
        });
    },
    // 修改课时
    updateVideo(video) {
      videoApi.updateVideo(video).then(response => {
        this.dialogVideoFormVisible = false;
        this.$message({
          type: "success",
          message: "修改课时成功!"
        });
        this.getChapterVideo();
      });
    },
    // 添加课时
    addVideo(video) {
      videoApi.addVideo(video).then(response => {
        this.dialogVideoFormVisible = false;
        this.$message({
          type: "success",
          message: "添加课时成功!"
        });
        this.getChapterVideo();
      });
    },
    // 添加或更新video
    saveOrUpdateVideo() {
     if (!this.video.id) {
        this.addVideo(this.video);
      } else {
        this.updateVideo(this.video);
      }
    },
    // 添加小节弹窗
    openVideoDialog(chapterId) {
      this.dialogVideoFormVisible = true;
      this.video.title = "";
      this.video.sort = 0;
      this.video.free = "";
      this.video.chapter_id = chapterId;
    },
    // 删除章节
    removeChapter(id) {
      this.$confirm("此操作将删除章节, 是否继续?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning"
      })
        .then(() => {
          chapterApi.deleteChapter(id).then(response => {
            this.$message({
              type: "success",
              message: "删除章节信息成功!"
            });
            this.getChapterVideo();
          });
        })
        .catch(() => {
          this.$message({
            type: "info",
            message: "已取消删除"
          });
        });
    },
    openEditChapterDialog(id) {
      chapterApi.getChapter(id).then(response => {
        this.chapter = response.data.chapter;
        this.dialogCHapterFormVisible = true;
      });
    },
    openChapterDialog() {
      this.chapter.id = "";
      this.chapter.title = "";
      this.chapter.sort = 0;
      this.dialogCHapterFormVisible = true;
    },
    saveOrUpdate() {
      if (!this.chapter.id) {
        this.createChapter(this.chapter);
      } else {
        this.updateChapter(this.chapter);
      }
    },
    updateChapter(chapter) {
      chapterApi.updateChapter(chapter).then(response => {
        this.dialogCHapterFormVisible = false;
        this.$message({
          type: "success",
          message: "修改课程信息成功!"
        });
        this.getChapterVideo();
      });
    },
    createChapter(chapter) {
      chapterApi.addChapter(chapter).then(response => {
        this.dialogCHapterFormVisible = false;
        this.$message({
          type: "success",
          message: "添加课程信息成功!"
        });
        this.getChapterVideo();
      });
    },
    previous() {
      this.$router.push({ path: `/course/info/${this.courseId}` });
    },
    next() {
      this.$router.push({ path: `/course/publish/${this.courseId}` });
    },
    getChapterVideo() {
      chapterApi.getChapterVideo(this.courseId).then(response => {
        this.chapterVideoList = response.data.allChapterVideo;
      });
    }
  }
};
</script>

<style scoped>
.chanpterList {
  position: relative;
  list-style: none;
  margin: 0;
  padding: 0;
}
.chanpterList li {
  position: relative;
}
.chanpterList p {
  float: left;
  font-size: 20px;
  margin: 10px 0;
  padding: 10px;
  height: 70px;
  line-height: 50px;
  width: 100%;
  border: 1px solid #ddd;
}
.chanpterList .acts {
  float: right;
  font-size: 14px;
}

.videoList {
  padding-left: 50px;
}
.videoList p {
  float: left;
  font-size: 14px;
  margin: 10px 0;
  padding: 10px;
  height: 50px;
  line-height: 30px;
  width: 100%;
  border: 1px dotted #ddd;
}
</style>