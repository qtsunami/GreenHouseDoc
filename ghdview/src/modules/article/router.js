import Vue from 'vue'
import VueRouter from 'vue-router'

Vue.use(VueRouter)

const routes = [
    { path: '/', name: 'article', component: r => { require(['./list/List'], r) }, meta: { title: '文章列表' }},
    { path: '/detail', name: 'article-detail', component: r => { require(['./list/Detail'], r) }, meta: { title: '文章详情' }},
]

export default new VueRouter({
    routes: routes
})