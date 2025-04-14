import { createRouter, createWebHistory } from 'vue-router'
import PrincipalPage from '../pages/PrincipalPage.vue'
import SelectGameMode from '../pages/SelectGameMode.vue'
import AutoSolveDimensions from '../pages/AutoSolveDimensions.vue'
import PlayableModeSelectDimensions from '../pages/PlayableModeSelectDimensions.vue'
import SelectGamesSaved from '../pages/SelectGamesSaved.vue'

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
    path: '/AutoSolveDimensions',
    name: 'AutoSolveDimensions',
    component: AutoSolveDimensions
  },
  {
    path: '/PlayableDimensions',
    name: 'PlayableDimensions',
    component: PlayableModeSelectDimensions
  },
  {
    path: '/SelectGamesSaved',
    name: 'SelectGamesSaved',
    component: SelectGamesSaved
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router