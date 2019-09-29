<template>
    <div id="app">
        <el-container style="">
        <el-header style="">
            <v-header></v-header>
        </el-header>
        <el-container>
            <el-main>
            <router-view></router-view>
            </el-main>
        </el-container>
        </el-container>
    </div>
</template>

<script>
  import Header from '@/components/Header'
  export default {
    name: 'app',
    components: {
      'v-header': Header
    }
  };
</script>