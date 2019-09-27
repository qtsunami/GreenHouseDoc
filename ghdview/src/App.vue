<template>
  <div id="app">
    <el-container style="">
      <el-header style="height: 61px; text-align: right; font-size: 12px; border-bottom: solid 1px #e6e6e6">
        <v-header></v-header>
      </el-header>
      <el-container>
        <el-aside width="200px" style="border-right: 1px solid #e6e6e6;">
          <navmenu></navmenu>
        </el-aside>
        
        <el-main style="margin-left: 200px;">
          <el-table :data="tableData">
            <el-table-column prop="date" label="日期" width="140">
            </el-table-column>
            <el-table-column prop="name" label="姓名" width="120">
            </el-table-column>
            <el-table-column prop="address" label="地址">
            </el-table-column>
          </el-table>
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>
<style>
  *{
    margin: 0;
    padding: 0;
  }
  html,body{
    width: 100%;
    height: 100%;
  }
  #app {
    width: 100%;
    height: 100%;
  }

  .el-header {
    background-color: #FFF;
    color: #333;
    line-height: 61px;
  }
  
  .el-aside {
    /* color: #333;
    height: 100vh; */
        color: #333;
      position: fixed;
      bottom: 0;
      top: 61px;
  }
</style>

<script>
  import NavMenu from '@/components/NavMenu'
  import Header from '@/components/Header'
  export default {
    name: 'app',
    data() {
      const item = {
        date: '2016-05-02',
        name: '王小虎',
        address: '上海市普陀区金沙江路 1518 弄'
      };
      return {
        tableData: Array(10).fill(item)
      }
    },
    components: {
      'navmenu': NavMenu,
      'v-header': Header
    }
  };
</script>