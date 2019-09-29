import Vue from 'vue'
import VueRouter from 'vue-router'

Vue.use(VueRouter)

const routes = [
    { path: '/', name: 'login', component: r => { require(['./login/Login'], r) }, meta: { title: 'console 登录' }},
    { path: '/register', name: 'login', component: r => { require(['./register/Register'], r) }, meta: { title: 'GreenHouseDoc 注册' }}
]

export default new VueRouter({
    routes: routes
})