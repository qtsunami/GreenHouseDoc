<template>
    <el-container id="app">
        <el-main class="main-content">
            <h1 style="margin-bottom: 20px;">文章列表</h1>
            
            <div class="box-card">
                <div class="title">
                    <router-link class="first-title" to="detail">没什么了不起的</router-link>
                    <el-text class="second-title">作者：Rosa</el-text>
                    <el-text class="second-title">日期：2019-04-12 12:10:11</el-text>
                </div>
                <div class="text item">
                    没什么了不起的没什么了不起的没什么了不起的没什么了不起的没什么了不起的没什么了不起的没什么了不起的没什么了不起的没什么了不起的没什么了不起的没什么了不起的没什么了不起的没什么了不起的没什么了不起的没什么了不起的没什么了不起的没什么了不起的没什么了不起的没什么了不起的没什么了不起的没什么了不起的没什么了不起的没什么了不起的没什么了不起的没什么了不起的没什么了不起的没什么了不起的没什么了不起的没什么了不起的没什么了不起的没什么了不起的没什么了不起的没什么了不起的没什么了不起的没什么了不起的没什么了不起的没什么了不起的没什么了不起的没什么了...
                </div>
                <el-link :underline="false" href="https://github.com/">查看全文</el-link>
            </div>

            <div class="box-card">
                <div class="title">
                    <router-link class="first-title" to="detail">没什么了不起的</router-link>
                    <el-text class="second-title">作者：Rosa</el-text>
                    <el-text class="second-title">日期：2019-04-12 12:10:11</el-text>
                </div>
                <div class="text item">
                    没什么了不起的没什么了不起的没什么了不起的没什么了不起的没什么了不起的
                </div>
                <el-link :underline="false" href="https://github.com/">查看全文</el-link>
            </div>

            <div class="box-card">
                <div class="title">
                    <router-link class="first-title" to="detail">没什么了不起的</router-link>
                    <el-text class="second-title">作者：Rosa</el-text>
                    <el-text class="second-title">日期：2019-04-12 12:10:11</el-text>
                </div>
                <div class="text item">
                    没什么了不起的没什么了不起的没什么了不起的没什么了不起的没什么了不起的
                </div>
                <el-link :underline="false" href="https://github.com/">查看全文</el-link>
            </div>

            <div class="box-card">
                <div class="title">
                    <router-link class="first-title" to="detail">没什么了不起的</router-link>
                    <el-text class="second-title">作者：Rosa</el-text>
                    <el-text class="second-title">日期：2019-04-12 12:10:11</el-text>
                </div>
                <div class="text item">
                    没什么了不起的没什么了不起的没什么了不起的没什么了不起的没什么了不起的
                </div>
                <router-link class="el-link el-link--default" to="detail">点击查看全文</router-link>
            </div>
            <div style="margin-top: 20px;">
                <el-pagination
                background
                layout="prev, pager, next"
                :total="1000">
                </el-pagination>
            </div>
            

        </el-main>

        

    </el-container>
    
</template>

<script>
export default {
    data(){
        return {
           author: "Rolly"
        };
    },
    methods: {
        
    }
};
</script>

<style scoped>
.el-container {
    width: 100%;
}
.main-content {
    margin-left: 50px;
}
.title .first-title{
    color: #409EFF; 
    font-size: 18px; 
}

.title .second-title{
    margin-left: 15px; 
    color: #838b8c; 
    font-size: 12px;
}


.box-card {
    margin-top: 20px;
    padding: 15px;
    border-radius: 5px;
    /* border: 1px dashed #ccc; */
}

.box-card:hover {
    background-color: #efefef;
}
.text {
font-size: 14px;
}

.item {
padding: 18px 0;
}

  
</style>
