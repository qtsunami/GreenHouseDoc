import Vue from 'vue'
import VueRouter from 'vue-router'

Vue.use(VueRouter)

const routes = [
    { path: '/', name: 'index', component: r => { require(['./console/Console'], r) }, meta: { title: '我的工作台' }}
]

export default new VueRouter({
    routes: routes
})