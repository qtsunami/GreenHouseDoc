<template>
  <div id="app">
    <el-container style="">
      <el-header>
        <v-header></v-header>
      </el-header>
      <el-container>
        <el-aside width="200px" style="border-right: 1px solid #e6e6e6;">
          <navmenu></navmenu>
        </el-aside>
        
        <el-main style="margin-left: 200px;">
          <el-breadcrumb separator-class="el-icon-arrow-right">
              <el-breadcrumb-item :to="{ path: '/' }">首页</el-breadcrumb-item>
              <el-breadcrumb-item>活动管理</el-breadcrumb-item>
              <el-breadcrumb-item>活动列表</el-breadcrumb-item>
              <el-breadcrumb-item>活动详情</el-breadcrumb-item>
          </el-breadcrumb>
          <router-view></router-view>
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>

<script>
  import NavMenu from '@/components/NavMenu'
  import Header from '@/components/Header'
  export default {
    name: 'app',
    components: {
      'navmenu': NavMenu,
      'v-header': Header
    }
  };
</script>