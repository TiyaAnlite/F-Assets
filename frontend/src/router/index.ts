import { createRouter, createWebHashHistory } from 'vue-router'
import AssetPage from '@/views/AssetPage.vue'
import RecordsPage from '@/views/RecordsPage.vue'

const router = createRouter({
  history: createWebHashHistory(),
  routes: [
    {
      path: '/',
      component: AssetPage,
    },
    {
      path: '/records/:id',
      component: RecordsPage,
      props: true,
    },
  ],
})

export default router
