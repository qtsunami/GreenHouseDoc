import Vue from 'vue'
import VueRouter from 'vue-router'

Vue.use(VueRouter)

const routes = [
    { path: '/', name: 'index', component: r => { require(['./workbench/Workbench'], r) }, meta: { title: 'console 登录' }},
    { path: '/workbench', name: 'index', component: r => { require(['./workbench/Workbench'], r) }, meta: { title: 'console 登录' }}
]

export default new VueRouter({
    routes: routes
})