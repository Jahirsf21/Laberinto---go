<script setup>
import { useRouter } from 'vue-router'
import { ref } from 'vue'
import { DeleteStartPoint, SetStartPointRandom } from '../../wailsjs/go/main/App'

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
    updatedMaze = await SetStartPointRandom(maze.value) 
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
    <div class="maze-grid">
      <div v-for="(row, rowIndex) in maze" :key="rowIndex" class="maze-row">
        <div v-for="(cell, colIndex) in row" :key="colIndex" class="maze-cell"
          :class="{
            'wall': cell === 3,
            'end': cell === 2,
          }">
            <span v-if="cell === 2">F</span>
        </div>
      </div>
    </div>
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

.maze-grid {
  display: inline-block;
  margin: 1rem auto;
  border: 2px solid #2c3e50;
  background: white;
}

.maze-row {
  display: flex;
}

.maze-cell {
  width: 30px;
  height: 30px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: bold;
  border: 1px solid #eee;
  position: relative;
}

.wall {
  background-color: #2c3e50;
}

.end {
  background-color: #ff6b6b;
  color: white;
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