import { createRouter, createWebHistory } from 'vue-router'
import PrincipalPage from '../pages/PrincipalPage.vue'
import SelectGameMode from '../pages/SelectGameMode.vue'
import AutoSolveDimensions from '../pages/AutoSolveDimensions.vue'
import PlayableModeSelectDimensions from '../pages/PlayableModeSelectDimensions.vue'
import SelectGamesSaved from '../pages/SelectGamesSaved.vue'
import SelectModeMazeSaved from '../pages/SelectModeMazeSaved.vue'
import Game from '../pages/Game.vue'

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
  },
  {
    path: '/SelectModeMazeSaved',
    name: 'SelectModeMazeSaved', 
    component: SelectModeMazeSaved
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