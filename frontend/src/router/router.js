import { createRouter, createWebHistory } from 'vue-router'
import PrincipalPage from '../pages/PrincipalPage.vue'
import SecondPage from '../pages/SecondPage.vue'
import ThirdPage from '../pages/ThirdPage.vue'
import Game from '../pages/Game.vue'

const routes = [
  {
    path: '/',
    name: 'Home', 
    component: PrincipalPage
  },
  {
    path: '/SecondPage',
    name: 'SecondPage',
    component: SecondPage
  },
  {
    path: '/ThirdPage',
    name: 'ThirdPage',
    component: ThirdPage
  },
  {
    path: '/game',
    name: 'Game',
    component: Game
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router