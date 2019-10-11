<template>
    <div id="app" v-wechat-title="$route.meta.title">
        <el-container style="">
			<el-header style="">
				<v-header></v-header>
			</el-header>
      <el-carousel :interval="4000" height="400px">
          <el-carousel-item v-for="item in 6" :key="item">
              <h3 class="medium">{{ item }}</h3>
          </el-carousel-item>
      </el-carousel>
			<el-main class="main-container">
				<router-view></router-view>
			</el-main>
			<el-footer>
				<v-footer></v-footer>
			</el-footer>
		</el-container>
    </div>
</template>

<script>
  import Header from '@/components/Header'
  import Footer from '@/components/Footer'
  export default {
    name: 'app',
    components: {
	  'v-header': Header,
	  'v-footer': Footer
    }
  };
</script>