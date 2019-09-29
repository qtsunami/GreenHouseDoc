import Vue from 'vue'
import VueRouter from 'vue-router'

Vue.use(VueRouter)

const routes = [
    { path: '/', name: 'project', component: r => { require(['./list/List'], r) }, meta: { title: '项目列表' }}
]

export default new VueRouter({
    routes: routes
})