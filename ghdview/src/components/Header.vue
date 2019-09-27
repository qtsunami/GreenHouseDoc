<template>
    <el-container>
        <h1 style="color: #da5fb1; font-size: 20px; margin-right: 100px;">GreenHouseDoc</h1>
        <el-menu :default-active="activeIndex" class="el-menu-demo" mode="horizontal" @select="handleSelect">
            <el-menu-item index="1">首页</el-menu-item>
            <el-menu-item index="2">标签</el-menu-item>
            <el-menu-item index="3">我的工作台</el-menu-item>
        </el-menu>
        <span style="position: absolute; font-size: 14px; right: 50px;">欢迎您，<b>王小虎</b></span>
    </el-container>
</template>

<style scoped>

</style>

<script>
    export default {
        data() {
        return {
            activeIndex: '1',
            activeIndex2: '1'
        };
        },
        methods: {
            handleSelect(key, keyPath) {
                // console.log(key, keyPath);
            }
        }
    }
</script>