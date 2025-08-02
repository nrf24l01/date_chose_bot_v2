import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/date',
      name: 'DateChoose',
      component: () => import('@/views/DateChose.vue')
    },
    {
      path: "/",
      name: "Home",
      redirect: { name: 'DateChoose' }
    }
  ],
})

export default router
