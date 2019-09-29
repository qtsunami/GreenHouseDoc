import Vue from 'vue'

import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';
// import '../../assets/css/style.css'  

import App from './App.vue'
import router from './router'

Vue.use(require('vue-wechat-title'))
Vue.use(ElementUI, { size: 'small', zIndex: 3000 });

new Vue({
    router,
    render: h => h(App)
}).$mount('#app')