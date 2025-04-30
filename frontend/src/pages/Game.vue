<script setup>
// Imports
import { useRouter } from 'vue-router'
import { ref, watch, onMounted, onUnmounted } from 'vue'
import { 
  SetStartPoint, 
  DeleteStartPoint, 
  SaveMaze, 
  GetPaths, 
  GetBestPath, 
  GetWorstPath, 
  MoveUp,
  MoveDown,
  MoveRight,
  MoveLeft
} from '../../wailsjs/go/main/App'

onMounted(() => {
  window.addEventListener('keydown', handleKeyDown);
});

// Router and session storage
const router = useRouter()
const savedData = JSON.parse(sessionStorage.getItem('gameData') || 'null')

// Data
const maze = ref(savedData?.maze || [])
const mode = ref(savedData?.mode || false)
const paths = ref([])
const bestPath = ref([])
const worstPath = ref([])
const selectedPath = ref(0)


//Game status
const IsSaved = ref(savedData?.IsSaved || false)
const hasStartPoint = ref(false)

//UI Controls
const selectingStartPoint = ref(false)
const showingSolution = ref(false)
const showingBestPath = ref(false)
const showingWorstPath = ref(false)

//Messages
const errorMessage = ref('')
const saveMessage = ref('')

// Watchers
watch(maze, (newMaze) => {
  hasStartPoint.value = newMaze.some(row => row.includes(1))
}, { immediate: true })

// Game navigation functions
const goHome = () => {
  router.push({ name: 'Home' })
}

// Start point management
const enableStartPointSelection = () => {
  selectingStartPoint.value = true
}

const selectCell = async (row, col) => {
  if (!selectingStartPoint.value) return
  if (maze.value[row][col] !== 0) {
    errorMessage.value = 'Solo puedes seleccionar una celda vacía'
    setTimeout(() => errorMessage.value = '', 3000)
    return
  }
  const updatedMaze = await SetStartPoint(maze.value, row, col)
  maze.value = updatedMaze
  const startX = maze.value.findIndex(row => row.includes(1))
  const startY = maze.value[startX]?.indexOf(1)
  worstPath.value = await GetWorstPath(maze.value, startX, startY)
  bestPath.value = await GetBestPath(maze.value, startX, startY)
  paths.value = await GetPaths(maze.value, startX, startY)
  selectingStartPoint.value = false
}

const handleKeyDown = async (event) => {
  let updatedMaze
  switch (event.key) {
    case 'w':
      updatedMaze = await MoveUp(maze.value)
      maze.value = updatedMaze
      break
    case 's':
      updatedMaze = await MoveDown(maze.value)
      maze.value = updatedMaze
      break
    case 'd':
      updatedMaze = await MoveRight(maze.value)
      maze.value = updatedMaze
      break
    case 'a':
      updatedMaze = await MoveLeft(maze.value)
      maze.value = updatedMaze
      break

  }

}

const deleteStartPoint = async () => {
  const updatedMaze = await DeleteStartPoint(maze.value)
  showingBestPath.value = false
  showingWorstPath.value = false
  showingSolution.value = false
  bestPath.value = []
  worstPath.value = []
  paths.value = []
  
  maze.value = updatedMaze
}

// Path display functions
const isPathCell = (row, col) => {
  if (!showingSolution.value) return false
  const path = paths.value[selectedPath.value]
  return path?.some(coord => coord.fila === row && coord.col === col)
}

const isBestPathCell = (row, col) => {
  if (!showingBestPath.value) return false
  const path = bestPath.value
  return path?.some(coord => coord.fila === row && coord.col === col)
}

const isWorstPathCell = (row, col) => {
  if (!showingWorstPath.value) return false
  const path = worstPath.value
  return path?.some(coord => coord.fila === row && coord.col === col)
}

// Path navigation controls
const showNextPath = () => {
  if (selectedPath.value < paths.value.length - 1) {
    selectedPath.value++
  }
}

const showPrevPath = () => {
  if (selectedPath.value >= 0) {
    selectedPath.value--
  }
}

// Solution display functions
const toggleShowPaths = async () => {
  if (selectingStartPoint.value) {
    selectingStartPoint.value = false
  }
  showingBestPath.value = false
  showingWorstPath.value = false
  
  if (!hasStartPoint.value) {
    errorMessage.value = 'No puedes ver los caminos sin un punto de inicio'
    setTimeout(() => errorMessage.value = '', 3000)
    return
  }

  if (mode.value && !paths.value.length) {
    const startX = maze.value.findIndex(row => row.includes(1))
    const startY = maze.value[startX]?.indexOf(1)
    paths.value = await GetPaths(maze.value, startX, startY)
    selectedPath.value = 0
  }
  
  showingSolution.value = !showingSolution.value
}

const showBestPath = async () => {
  showingSolution.value = false
  showingWorstPath.value = false
  
  if (!hasStartPoint.value) {
    errorMessage.value = 'No puedes ver la solución sin un punto de inicio'
    setTimeout(() => errorMessage.value = '', 3000)
    return
  }
  
  showingBestPath.value = !showingBestPath.value
}

const showWorstPath = async () => {
  showingSolution.value = false
  showingBestPath.value = false
  
  if (!hasStartPoint.value) {
    errorMessage.value = 'No puedes ver la solución sin un punto de inicio'
    setTimeout(() => errorMessage.value = '', 3000)
    return
  }
  
  showingWorstPath.value = !showingWorstPath.value
}

// Game saving
const saveMaze = async () => {
  if (selectingStartPoint.value) {
    selectingStartPoint.value = false
  }
  
  const mazeName = prompt("Ingrese el nombre para el laberinto:")
  if (!mazeName) {
    saveMessage.value = 'Debe ingresar un nombre para el laberinto'
    setTimeout(() => saveMessage.value = '', 3000)
    return
  }

  try {
    await SaveMaze(maze.value, mazeName, mode.value)
    saveMessage.value = 'Laberinto guardado correctamente'
    IsSaved.value = true
    setTimeout(() => saveMessage.value = '', 3000)
  } catch (error) {
    saveMessage.value = 'Error al guardar el laberinto'
    setTimeout(() => saveMessage.value = '', 3000)
    console.error(error)
  }
}
</script>

<template>
  <div class="game-container">
    <h1>Maze</h1>
    
    <div class="maze-grid">
      <div v-for="(row, rowIndex) in maze" :key="rowIndex" class="maze-row">
        <div v-for="(cell, colIndex) in row" :key="colIndex" class="maze-cell"
          :class="{
            'wall': cell === 3,
            'path': cell === 0,
            'start': cell === 1,
            'end': cell === 2,
            'player' : cell === 4,
            'solution-path': showingSolution && isPathCell(rowIndex, colIndex),
            'best-path': showingBestPath && isBestPathCell(rowIndex, colIndex),
            'worst-path': showingWorstPath && isWorstPathCell(rowIndex, colIndex)
          }"
          @click="selectCell(rowIndex, colIndex)" 
        >
          <span v-if="cell === 1">I</span>
          <span v-else-if="cell === 2">F</span>
        </div>
      </div>
    </div>

    <div v-if="showingSolution && paths.length" class="path-navigation">
      <button @click="showPrevPath" :disabled="selectedPath === 0" class="nav-button">
        ← Anterior
      </button>
      <span class="path-counter">
        Camino {{ selectedPath + 1 }} de {{ paths.length }}
      </span>
      <button @click="showNextPath" :disabled="selectedPath === paths.length - 1" class="nav-button">
        Siguiente →
      </button>
    </div>

    <!-- NUEVA ORGANIZACIÓN DE BOTONES -->
    <div class="controls-group">
      <!-- Primer grupo -->
      <div class="controls">
        <button v-if="mode && !hasStartPoint" @click="enableStartPointSelection" class="start-button">
          Establecer punto de inicio
        </button> 
        <button v-if="mode && hasStartPoint" @click="deleteStartPoint" class="delete-button">
          Borrar punto de inicio
        </button>

        <button @click="showBestPath" class="best-path-button">
          {{ showingBestPath ? "Ocultar" : "Mostrar solución" }}
        </button>
        <button v-if="worstPath.length > bestPath.length" @click="showWorstPath" class="worst-path-button">
          {{ showingWorstPath ? "Ocultar" : "Mostrar peor camino" }}
        </button>

        <button @click="toggleShowPaths" class="solution-button">
          {{ showingSolution ? 'Ocultar caminos' : 'Mostrar caminos' }}
        </button>
      </div>

      <!-- Segundo grupo -->
      <div class="controls">
        <button v-if="!IsSaved" @click="saveMaze" class="save-button">
          Guardar Laberinto
        </button>
        <button @click="goHome" class="back-button">
          Terminar juego
        </button>
      </div>
    </div>

  </div>

  <div v-if="errorMessage" class="message error-message">
    {{ errorMessage }}
  </div>
</template>

<style scoped>

.game-container {
  max-width: 100%;
  margin: 0 auto;
  padding: 1rem;
  text-align: center;
}

h1 {
  color: #42b883;
  margin-bottom: 1rem;
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
  background-image: url('@/assets/pared.png');
  background-size: contain;
  background-position: center;
  background-repeat: no-repeat;
  color: transparent;
  font-size: 0; 
}

.path {
  background-color: #f8f9fa;
}

.start {
  background-color: #42b883;
  color: white;
}

.end {
  background-image: url('@/assets/meta.png');
  background-size: contain;
  background-position: center;
  background-repeat: no-repeat;
  color: transparent;
  font-size: 0;
}

.player {
  background-image: url('@/assets/character.png');
  background-size: contain;
  background-position: center;
  background-repeat: no-repeat;
  color: transparent;
  font-size: 5;
}

.solution-path {
  background-color: rgba(52, 152, 219, 0.5);
}

.best-path {
  background-color: #42b883;
}

.worst-path {
  background-color: #e74c3c;
}


.controls-group {
  display: flex;
  flex-direction: column;
  gap: 2rem; 
  margin-top: 2rem;
}

.controls {
  display: flex;
  justify-content: center;
  flex-wrap: wrap;
  gap: 1rem;
}

.path-navigation {
  margin: 1.5rem 0;
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
  align-items: center;
  gap: 1rem;
  padding: 0.8rem;
  background-color: #2c3e50;
  border-radius: 8px;
}

.path-counter {
  font-weight: bold;
  min-width: 150px;
}


button {
  padding: 0.8rem 1.5rem;
  border: none;
  border-radius: 6px;
  font-size: 1rem;
  cursor: pointer;
  transition: all 0.3s ease;
}


.nav-button {
  padding: 0.6rem 1rem;
  background-color: #3498db;
  color: white;
}

.start-button {
  background-color: #3498db;
  color: white;
}

.delete-button {
  background-color: #3498db;
  color: white;
}

.solution-button {
  background-color: #3498db;
  color: white;
}

.back-button {
  background-color: #3498db;
  color: white;
}

.save-button {
  background-color: #3498db;
  color: white;
}

.best-path-button {
  padding: 0.6rem 1rem;
  background-color: #3498db;
  color: white;
}

.worst-path-button {
  padding: 0.6rem 1rem;
  background-color: #3498db;
  color: white;
}


button:hover,
.nav-button:hover:not(:disabled),
.best-path-button:hover,
.worst-path-button:hover,
.start-button:hover,
.delete-button:hover,
.solution-button:hover,
.back-button:hover,
.save-button:hover {
  opacity: 0.9;
  transform: translateY(-2px);
}

.nav-button:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}


.message {
  margin-top: 1rem;
  padding: 0.8rem;
  border-radius: 6px;
  display: inline-block;
}

.error-message {
  background-color: #e74c3c;
  color: white;
}

.save-message {
  background-color: #00ff44;
  color: white;
}


@media (max-width: 768px) {
  .controls, 
  .path-navigation {
    flex-direction: column;
    align-items: center;
  }
  
  button,
  .nav-button,
  .best-path-button {
    width: 100%;
    max-width: 250px;
    margin-bottom: 0.5rem;
  }
  
  .maze-cell {
    width: 25px;
    height: 25px;
    font-size: 0.8rem;
  }
  
  .path-navigation {
    gap: 0.5rem;
  }
  
  .path-counter {
    order: -1;
    width: 100%;
    margin-bottom: 0.5rem;
  }
}
</style>
