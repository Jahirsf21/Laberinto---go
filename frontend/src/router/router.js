import { createRouter, createWebHistory } from 'vue-router'
import PrincipalPage from '../pages/PrincipalPage.vue'
import SelectGameMode from '../pages/SelectGameMode.vue'
import AutoSolveDimensions from '../pages/AutoSolveDimensions.vue'
import PlayableModeSelectDimensions from '../pages/PlayableModeSelectDimensions.vue'

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
  },
  {
    path: '/SelectGameMode/AutoSolveDimensions',
    name: 'AutoSolveDimensions',
    component: AutoSolveDimensions
  },
  {
    path: '/SelectGameMode/PlayableDimensions',
    name: 'PlayableDimensions',
    component: PlayableModeSelectDimensions
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router