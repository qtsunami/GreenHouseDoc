<template>
    <el-container>
        <h1 style="color: #da5fb1; font-size: 20px; margin-right: 100px;">GreenHouseDoc</h1>
        <el-menu :default-active="activeIndex" class="el-menu-demo" mode="horizontal" @select="handleSelect">
            <el-menu-item index="1"><a href="/index.html">首页</a></el-menu-item>
            <el-menu-item index="2"><a href="/project.html">项目</a></el-menu-item>
            <el-menu-item index="3"><a href="/article.html">文章</a></el-menu-item>
            <el-menu-item index="4"><a href="/booknote.html">读书笔记</a></el-menu-item>
            <el-menu-item index="5"><a href="/workbench.html">我的工作台</a></el-menu-item>
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
                topNav: {
                    'index': '1',
                    'project': '2',
                    'article': '3',
                    'booknote': '4',
                    'workbench': '5'
                }
            };
        },
        computed:{
        },
        created() {
            let pathname = location.pathname;
            let nav = location.pathname.substring(1, pathname.lastIndexOf('.'));
            this.activeIndex = eval("this.topNav." + nav)
        },
        methods: {
            handleSelect(key, keyPath) {
                // console.log(key, keyPath);
            }
        }
    }
</script>