import Vue from 'vue';
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';
import './assets/css/style.css'  
import router from './router'
import App from './App.vue';

Vue.config.productionTip = false

Vue.use(ElementUI, { size: 'small', zIndex: 3000 });

new Vue({
  router,
  render: h => h(App),
}).$mount('#app')
