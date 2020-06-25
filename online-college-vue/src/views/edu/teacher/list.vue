<template>
  <div class="app-container">
    <el-form :inline="true" class="demo-form-inline">
      <el-form-item>
        <el-input v-model="teacherQuery.name" placeholder="讲师名称"></el-input>
      </el-form-item>
      <el-form-item>
        <el-select v-model="teacherQuery.level" clearable placeholder="讲师头衔">
          <el-option label="首席讲师" :value="2"></el-option>
          <el-option label="高级讲师" :value="1"></el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="添加时间">
        <el-date-picker
          v-model="teacherQuery.begin"
          placeholder="开始时间"
          type="datetime"
          value-format="yyyy-MM-dd HH:mm:ss"
          default-time="00:00:00"
        ></el-date-picker>
        <el-date-picker
          v-model="teacherQuery.end"
          placeholder="结束时间"
          type="datetime"
          value-format="yyyy-MM-dd HH:mm:ss"
          default-time="00:00:00"
        ></el-date-picker>
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
      <el-table-column label="序号" width="70" align="center">
        <template slot-scope="scope">{{(page - 1)*limit + scope.$index+1}}</template>
      </el-table-column>
      <el-table-column prop="name" label="姓名" width="80"></el-table-column>
      <el-table-column label="讲师头衔" width="80">
        <template slot-scope="scope">{{scope.row.level==1?'高级讲师':'首席讲师'}}</template>
      </el-table-column>
      <el-table-column prop="career" label="讲师资历"></el-table-column>
      <!-- <el-table-column prop="intro" label="讲师简介"></el-table-column> -->
      <el-table-column prop="gmtCreate" label="添加时间"></el-table-column>
      <el-table-column prop="sort" label="排序"></el-table-column>
      <el-table-column label="操作" width="200" align="center">
        <template slot-scope="scope">
          <router-link :to="'/teacher/edit/'+scope.row.id">
            <el-button type="primary" size="mini" icon="el-icon-edit">修改</el-button>
          </router-link>
          <el-button
            type="danger"
            size="mini"
            icon="el-icon-delete"
            @click="removeTeacherById(scope.row.id)"
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
import teacher from "@/api/edu/teacher";

export default {
  data() {
    return {
      page: 1,
      limit: 5,
      teacherQuery: {},
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
      teacher
        .getTeacherListPage(this.page, this.limit, this.teacherQuery)
        .then(response => {
          this.list = response.data.rows;
          this.total = response.data.total;
          this.listLoading = false;
        })
        .catch(error => {
          this.$message.error("查询失败");
        });
    },
    resetDate() {
      this.teacherQuery = {};
      this.getList();
    },
    removeTeacherById(id) {
      this.$confirm("此操作将永久删除讲师记录, 是否继续?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning"
      })
        .then(() => {
          teacher
            .removeTeacherById(id)
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
