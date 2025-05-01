<script setup>
import { ref, onMounted } from 'vue'
import { GetMazes, DeleteStartPoint, SetStartPointRandom} from '../../wailsjs/go/main/App'
import { useRouter } from 'vue-router'

const router = useRouter()
const mazes = ref([])
const selectedMaze = ref(null)
const mode = ref(true)

// Functions
const loadMazes = async () => {
  const response = await GetMazes()
  mazes.value = response.Mazes || []
}

const selectMaze = (maze) => {
  selectedMaze.value = maze
}

const startGame = async () => {
  let updatedMaze = selectedMaze.value.matrix
  if (selectedMaze.value.mode === true) {
    if (mode.value === false) {
      updatedMaze = await SetStartPointRandom(updatedMaze)
    }
  } else {
    if (mode.value === true) {
      updatedMaze = await DeleteStartPoint(updatedMaze)
    } else {
      updatedMaze = await SetStartPointRandom(updatedMaze)
    }
  }
  
  sessionStorage.setItem('gameData', JSON.stringify({
    maze: updatedMaze,
    mode: mode.value,
    isSaved: true,
  }))

  router.push({ name: 'Game' })
}

const goHome = () => {
  router.push({ name: 'Home' })
}


onMounted(() => {
  loadMazes()
})
</script>

<template>
  <div class="container">
    <h1>Selecciona tu laberinto</h1>
    <p class="subtitle">Elige un laberinto guardado y configura el modo de juego</p>
    
    <div class="content-wrapper">
      <div class="maze-selection-panel">
        <div class="maze-buttons-grid">
          <div 
            v-for="maze in mazes" 
            :key="maze.name" 
            class="maze-button"
            :class="{ 'selected': selectedMaze?.name === maze.name }"
            @click="selectMaze(maze)"
          >
            <h3>{{ maze.name }}</h3>
            <p>Modo: {{ maze.mode ? 'Auto solución' : 'Jugable' }}</p>
            <p>Tamaño: {{ maze.matrix.length }}x{{ maze.matrix[0].length }}</p>
          </div>
        </div>
      </div>
      <div class="preview-controls-panel">
        <div v-if="selectedMaze" class="maze-preview">
          <div class="maze-grid">
            <div v-for="(row, rowIndex) in selectedMaze.matrix" :key="rowIndex" class="maze-row">
              <div v-for="(cell, colIndex) in row" :key="colIndex" class="maze-cell"
                :class="{
                  'wall': cell === 3,
                  'end': cell === 2,
                  'start': cell === 1
                }">
                <span v-if="cell === 1">I</span>
                <span v-else-if="cell === 2">F</span>
              </div>
            </div>
          </div>
        </div>
        <div v-else class="no-maze-selected">
          <p>Selecciona un laberinto para ver la vista previa</p>
        </div>
        <div class="controls-container">
          <div class="mode-section">
            <p class="subtitle">Selecciona el modo</p>
            <div class="mode-toggle-container">
              <span class="mode-label">Jugable</span>
              <label class="switch">
                <input type="checkbox" v-model="mode">
                <span class="slider round"></span>
              </label>
              <span class="mode-label">Auto Solución</span>
            </div>
          </div>
          <div class="action-buttons">
            <button 
              @click="startGame" 
              class="start-button"
              :disabled="!selectedMaze"
            >
              Iniciar Juego
            </button>
            <button @click="goHome" class="back-button">
              Volver al Inicio
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 2rem;
  text-align: center;
}

h1 {
  color: #ffffff;
  margin-bottom: 0.5rem;
}

.subtitle {
  color: #7f8c8d;
  margin-bottom: 2rem;
  font-size: 1.1rem;
}

.content-wrapper {
  display: flex;
  gap: 2rem;
  margin-top: 2rem;
}

.maze-selection-panel {
  flex: 1;
  max-height: 600px;
  overflow-y: auto;
  padding: 1rem;
  background-color: #2c3e50;
  border-radius: 8px;
}

.preview-controls-panel {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 2rem;
}

.maze-buttons-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));
  gap: 1rem;
}

.maze-button {
  padding: 1.2rem;
  color: white;
  background-color: #42b883;
  border: none;
  border-radius: 8px;
  font-size: 1rem;
  cursor: pointer;
  transition: all 0.3s ease;
}

.maze-button:hover {
  transform: translateY(-3px);
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  background-color: #309e6d;
}


.maze-button h3 {
  margin: 0 0 0.3rem 0;
  font-size: 1rem;
  color: #ffffff;
}

.maze-button p {
  margin: 0.2rem 0;
  font-size: 0.8rem;
  color: #dfe6e9;
}

.maze-preview {
  flex: 1;
  display: flex;
  justify-content: center;
  align-items: center;
  background-color: #2c3e50;
  border-radius: 8px;
  padding: 1rem;
}

.no-maze-selected {
  flex: 1;
  display: flex;
  justify-content: center;
  align-items: center;
  background-color: #2c3e50;
  border-radius: 8px;
  color: #7f8c8d;
  padding: 2rem;
}

.maze-grid {
  display: inline-block;
  margin: 0 auto;
  border: 2px solid #34495e;
  background: white;
}

.maze-row {
  display: flex;
}

.maze-cell {
  width: 20px;
  height: 20px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: bold;
  border: 1px solid #eee;
  font-size: 0.7rem;
}

.wall {
  background-color: #2c3e50;
}

.end {
  background-color: #ff6b6b;
  color: white;
}

.start {
  background-color: #42b883;
  color: white;
}

.controls-container {
  display: flex;
  flex-direction: column;
  gap: 1.5rem;
  background-color: #2c3e50;
  border-radius: 8px;
  padding: 1.5rem;
}

.mode-section {
  margin: 1.5rem 0;
}

.mode-toggle-container {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 1rem;
  margin: 1rem 0;
}

.mode-label {
  color: #ffffff;
  font-size: 1.1rem;
  min-width: 100px;
}

.switch {
  position: relative;
  display: inline-block;
  width: 60px;
  height: 34px;
  flex-shrink: 0;
}

.switch input {
  opacity: 0;
  width: 0;
  height: 0;
}

.slider {
  position: absolute;
  cursor: pointer;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: #ccc;
  transition: .4s;
}

.slider:before {
  position: absolute;
  content: "";
  height: 26px;
  width: 26px;
  left: 4px;
  bottom: 4px;
  background-color: white;
  transition: .4s;
}

input:checked + .slider {
  background-color: #2196F3;
}

input:checked + .slider:before {
  transform: translateX(26px);
}

.slider.round {
  border-radius: 34px;
}

.slider.round:before {
  border-radius: 50%;
}


.mode-button:hover {
  transform: translateY(-3px);
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  background-color: #309e6d;
}


.action-buttons {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.start-button {
  padding: 1rem 2rem;
  background-color: #42b883;
  color: white;
  border: none;
  border-radius: 6px;
  font-size: 1rem;
  cursor: pointer;
  transition: all 0.3s ease;
}

.start-button:hover {
  transform: translateY(-3px);
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  background-color: #309e6d;
}

.start-button:disabled {
  background-color: #579e7e;
  cursor: not-allowed;
}

.back-button {
  padding: 1rem 2rem;
  background-color: #42b883;
  color: white;
  border: none;
  border-radius: 6px;
  font-size: 1rem;
  cursor: pointer;
  transition: all 0.3s ease;
}

.back-button:hover {
  transform: translateY(-3px);
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  background-color: #309e6d;
}


</style>