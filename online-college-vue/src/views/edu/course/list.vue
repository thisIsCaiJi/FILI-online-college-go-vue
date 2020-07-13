<template>
  <div class="app-container">
    <el-form :inline="true" class="demo-form-inline">
      <el-form-item>
        <el-input v-model="courseQuery.title" placeholder="课程名称"></el-input>
      </el-form-item>
      <el-form-item>
        <el-select v-model="courseQuery.status" clearable placeholder="课程状态">
          <el-option label="未发布" value="Draft"></el-option>
          <el-option label="已发布" value="Normal"></el-option>
        </el-select>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="getList()">查询</el-button>
        <el-button type="default" @click="resetDate()">重置</el-button>
      </el-form-item>
    </el-form>

    <el-table
      v-loading="listLoading"
      :data="list"
      element-loading-text="Loading"
      border
      fit
      highlight-current-row
    >
      <el-table-column label="序号" width="80px" align="center">
        <template slot-scope="scope">{{(page - 1)*limit + scope.$index+1}}</template>
      </el-table-column>
      <el-table-column prop="title" label="课程名称" ></el-table-column>
      <el-table-column label="课程状态" width="100px">
        <template slot-scope="scope">{{scope.row.status=='Normal'?'已发布':'未发布'}}</template>
      </el-table-column>
      <el-table-column label="操作" width="400px" align="center">
        <template slot-scope="scope">
          <router-link :to="'/course/info/'+scope.row.id">
            <el-button type="primary" size="mini" icon="el-icon-edit">修改课程基本信息</el-button>
          </router-link>
          <router-link :to="'/course/chapter/'+scope.row.id">
            <el-button type="primary" size="mini" icon="el-icon-edit">修改课程大纲</el-button>
          </router-link>
          <el-button
            type="danger"
            size="mini"
            icon="el-icon-delete"
            @click="removeCourseInfoById(scope.row.id)"
          >删除</el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-pagination
      :current-page="page"
      :page-size="limit"
      :total="total"
      layout="total,prev, pager, next,jumper"
      @current-change="getList"
      style="padding: 10px 0; text-align:center;"
    ></el-pagination>
  </div>
</template>
 
<script>
import courseApi from "@/api/edu/course";

export default {
  data() {
    return {
      page: 1,
      limit: 5,
      courseQuery: {status:'',title:''},
      total: 0,
      list: null,
      listLoading: true
    };
  },
  created() {
    this.getList();
  },
  methods: {
    getList(page = 1) {
      this.page = page;
      this.listLoading = true;
      courseApi
        .getCourseList(this.page, this.limit, this.courseQuery)
        .then(response => {
          this.list = response.data.rows;
          this.total = response.data.total;
          this.listLoading = false;
        })
    
    },
    resetDate() {
      this.courseQuery = {};
      this.getList();
    },
    removeCourseInfoById(id) {
      this.$confirm("此操作将永久删除课程, 是否继续?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning"
      })
        .then(() => {
          courseApi
            .removeCourseInfoById(id)
            .then(response => {
              this.$message({
                message: "删除成功",
                type: "success"
              });
              this.getList();
            })
            .catch(error => {
              this.$message.error("删除失败");
            });
        })
        .catch(() => {
          this.$message({
            type: "info",
            message: "已取消删除"
          });
        });
    }
  }
};
</script>
