<template>
	<div id="app">
		<el-container style="">
			<el-header style="">
				<v-header></v-header>
			</el-header>
			<el-main class="main-container">
				<router-view></router-view>
			</el-main>
			<el-footer class="el-footer">
				<v-footer></v-footer>
			</el-footer>
		</el-container>
	</div>
</template>

<style scoped>


</style>
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