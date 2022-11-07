import { createRouter, createWebHashHistory } from 'vue-router'

const routes = [{
  path: '/',
  redirect: '/login'
},
{
  path: '/init',
  name: 'Init',
  component: () => import('@/view/init/index.vue')
},
{
  path: '/login',
  name: 'Login',
  component: () => import('@/view/login/index.vue')
},
// {
//   path: '/zaloAuthentication',
//   name: 'zaloAuthentication',
//   component: () => import('@/view/zaloApplication/zaloApplicationAuth.vue')
// }
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

export default router
