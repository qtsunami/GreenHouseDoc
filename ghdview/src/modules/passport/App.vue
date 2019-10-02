<template>
    <div id="app" v-wechat-title="$route.meta.title">
        <el-header style="line-height: 60px; box-shadow: 0 0 4px #717171;">
            <h1 style="margin: 0 200px; color: #da5fb1; font-size: 20px;">GreenHouseDoc</h1>
        </el-header>
        <router-view></router-view>
        <el-footer class="el-footer">
            <Footer></Footer>
        </el-footer>
    </div>
</template>

<script>
    import Footer from '@/components/Footer'
    export default {
        name: 'app',
        components: {
            'Footer': Footer
        }
    };
</script>

<style scoped>

</style>