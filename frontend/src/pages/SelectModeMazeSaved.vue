<script setup>
import { useRouter } from 'vue-router'
import { ref } from 'vue'
import { DeleteStartPoint, SetSPointRand } from '../../wailsjs/go/main/App'

const router = useRouter()
const savedData = JSON.parse(sessionStorage.getItem('gameData') || 'null')

const maze = ref(savedData?.maze || [])
const mode = ref(savedData?.mode || false)

const goSelectGamesSaved = () => {
  router.push({ name: 'SelectGamesSaved' }) 
}

const selectAutoSolve = async () => {
  let updatedMaze = await DeleteStartPoint(maze.value)
  sessionStorage.setItem('gameData', JSON.stringify({
    maze: updatedMaze,
    mode: true, 
    IsSaved: true
  }))
  router.push({ name: 'Game' })
}

const selectPlayable = async () => {
  let updatedMaze = maze.value
  if (mode.value === true) {
    updatedMaze = await SetSPointRand(maze.value) 
  }
  sessionStorage.setItem('gameData', JSON.stringify({
    maze: updatedMaze,
    mode: false, 
    IsSaved: true
  }))
  router.push({ name: 'Game' })
}
</script>

<template>
  <div class="mode-selection-container">
    <h1>Selecciona el modo en el que quieres jugar</h1>
    <div class="buttons-container">
      <button @click="selectAutoSolve" class="mode-button">Auto solución</button>
      <button @click="selectPlayable" class="mode-button">Jugable</button>
    </div>
    <button @click="goSelectGamesSaved" class="back-button">← Volver</button>
  </div>
</template>

<style scoped>
.mode-selection-container {
  text-align: center;
  padding: 2rem;
}

.buttons-container {
  display: flex;
  justify-content: center;
  gap: 20px;
  margin: 20px 0;
}

.mode-button {
  padding: 12px 24px;
  background-color: #42b883;
  color: white;
  border: none;
  border-radius: 6px;
  font-size: 1rem;
  cursor: pointer;
  transition: background-color 0.3s;
  width: 200px;
}

.mode-button:hover {
  background-color: #058938;
}

.back-button {
  padding: 12px 24px;
  background-color: #f0f0f0;
  color: #2c3e50;
  border: none;
  border-radius: 6px;
  font-size: 1rem;
  cursor: pointer;
  transition: background-color 0.3s;
  width: 200px;
  margin-top: 20px;
}

.back-button:hover {
  background-color: #e0e0e0;
}
</style>