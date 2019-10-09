<template>
    <el-container id="app">
        <el-row>
            <el-col :span="5" v-for="(o, index) in 4" :key="o" :offset="index > 0 ? 1 : 0" class="col-line">
                <el-card :body-style="{ padding: '0px' }">
                <img src="https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png" class="image">
                <div style="padding: 14px;">
                    <span>GreenHouseDoc 文档</span>
                    <div class="bottom clearfix">
                    <time class="author">作者：{{ author }}</time>
                    <el-button type="text" class="button">操作按钮</el-button>
                    </div>
                </div>
                </el-card>
            </el-col>

            <el-col :span="5" v-for="(o, index) in 4" :key="o" :offset="index > 0 ? 1 : 0" class="col-line">
                <el-card :body-style="{ padding: '0px' }">
                <img src="https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png" class="image">
                <div style="padding: 14px;">
                    <span>好吃的汉堡</span>
                    <div class="bottom clearfix">
                    <time class="author">作者：{{ author }}</time>
                    <el-button type="text" class="button">操作按钮</el-button>
                    </div>
                </div>
                </el-card>
            </el-col>
            <el-col :span="12">
                <el-pagination
                background
                layout="prev, pager, next"
                :total="1000">
                </el-pagination>

            </el-col>

        
        </el-row>


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
.author {
    font-size: 13px;
    color: #999;
  }
.col-line {
    margin-bottom: 40px;
}
  
  .bottom {
    margin-top: 13px;
    line-height: 12px;
  }

  .button {
    padding: 0;
    float: right;
  }

  .image {
    width: 100%;
    display: block;
  }

  .clearfix:before,
  .clearfix:after {
      display: table;
      content: "";
  }
  
  .clearfix:after {
      clear: both
  }

</style>

