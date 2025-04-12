import { createRouter, createWebHistory } from 'vue-router'
import PrincipalPage from '../pages/PrincipalPage.vue'
import SelectGameMode from '../pages/SelectGameMode.vue'

const routes = [
  {
    path: '/',
    name: 'Home', 
    component: PrincipalPage
  },
  {
    path: '/SelectGameMode',
    name: 'SelectGameMode',
    component: SelectGameMode
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router