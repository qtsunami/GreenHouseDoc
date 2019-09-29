import Vue from 'vue'
import Router from 'vue-router'
import menus from '@/config/menu-config'

Vue.use(Router)

var routes = []

routes.push({
    path: `/`,
    name: '',
    component: () => import(`@/pages/index.vue`)
})

menus.forEach((item) => {
    item.sub.forEach((sub) => {
        routes.push({
            path: `/${sub.componentName}`,
            name: sub.componentName,
            component: () => import(`@/pages/${sub.componentName}`)
            // component: dashboard
        })
    })
})

console.log(routes)

export default new Router({ routes })