import { createRouter, createWebHistory } from 'vue-router'
import PrincipalPage from '../pages/PrincipalPage.vue'
import SecondPage from '../pages/SecondPage.vue'
import ThirdPage from '../pages/ThirdPage.vue'
import Game from '../pages/Game.vue'
import Result from '../pages/Result.vue'

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
    path: '/Game',
    name: 'Game',
    component: Game
  },
  {
    path: '/Result',
    name: 'Result',
    component: Result
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router