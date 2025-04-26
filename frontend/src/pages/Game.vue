<script setup>
import { useRouter } from 'vue-router'
import { ref, watch } from 'vue'
import { SetStartPoint, DeleteStartPoint, SaveMaze, GetPaths, GetBestPath } from '../../wailsjs/go/main/App'

const router = useRouter()
const savedData = JSON.parse(sessionStorage.getItem('gameData') || 'null')

const maze = ref(savedData?.maze || [])
const mode = ref(savedData?.mode || false)
const allPaths = ref([])
const bestPath = ref([])
const currentPathIndex = ref(0)

//States of the game
const IsSaved = ref(savedData?.IsSaved || false)
const hasStartPoint = ref(false) 

//States for buttons
const selectingStartPoint = ref(false)
const showingSolution = ref(false)

//Messages
const errorMessage = ref('')
const saveMessage = ref('') 

const showSolution = async() => {
  if (selectingStartPoint.value == true) {
    selectingStartPoint.value = false
  }
  if (hasStartPoint.value == false) {
    errorMessage.value = 'No puedes ver la solución sin un punto de inicio'
    setTimeout(() => errorMessage.value = '', 3000) 
    return
  }
  if (!allPaths.value.length) {
    const startX = maze.value.findIndex(row => row.includes(1))
    const startY = maze.value[startX]?.indexOf(1) 
    allPaths.value = await GetPaths(maze.value, startX, startY)
    bestPath.value = await GetBestPath(maze.value, startX, startY)
  }
  
  showingSolution.value = !showingSolution.value
}

const showNextPath = () => {
  if (currentPathIndex.value < allPaths.value.length - 1) {
    currentPathIndex.value++
  }
}

const showPrevPath = () => {
  if (currentPathIndex.value > 0) {
    currentPathIndex.value--
  }
}

const showBestPath = () => {
  currentPathIndex.value = allPaths.value.findIndex(path => 
    JSON.stringify(path) === JSON.stringify(bestPath.value)
  )
}

const isPathCell = (row, col) => {
  if (!allPaths.value.length || currentPathIndex.value >= allPaths.value.length) return false
  const currentPath = allPaths.value[currentPathIndex.value]
  return currentPath.some(coord => coord.fila === row && coord.col === col)
}

const isBestPathCell = (row, col) => {
  if (!bestPath.value.length) return false
  return bestPath.value.some(coord => coord.fila === row && coord.col === col)
}


watch(maze, (newMaze) => {
  hasStartPoint.value = newMaze.some(row => row.includes(1))
}, { immediate: true })

const goHome = () => {
  router.push({ name: 'Home' })
}

const enableStartPointSelection = () => {
  selectingStartPoint.value = true
}

const selectCell = async (row, col) => {
  if (!selectingStartPoint.value) 
    return 
  if (maze.value[row][col] !== 0) {
    errorMessage.value = 'Solo puedes seleccionar una celda vacía'
    setTimeout(() => errorMessage.value = '', 3000) 
    return
  }
  const updatedMaze = await SetStartPoint(maze.value, row, col)
  maze.value = updatedMaze 
  allPaths.value = []
  selectingStartPoint.value = false
}

const deleteStartPoint = async () => {
  const updatedMaze = await DeleteStartPoint(maze.value)
  if (showingSolution.value) {
    showingSolution.value = false
  }
  maze.value = updatedMaze 
}

const saveMaze = async () => {
  if (selectingStartPoint.value == true) {
    selectingStartPoint.value = false
  }
  const mazeName = prompt("Ingrese el nombre para el laberinto:");
  if (!mazeName) {
    saveMessage.value = 'Debe ingresar un nombre para el laberinto';
    setTimeout(() => saveMessage.value = '', 3000);
    return;
  }

  try {
    await SaveMaze(maze.value, mazeName, mode.value);
    saveMessage.value = 'Laberinto guardado correctamente';
    IsSaved.value = true;
    setTimeout(() => saveMessage.value = '', 3000);
  } catch (error) {
    saveMessage.value = 'Error al guardar el laberinto';
    setTimeout(() => saveMessage.value = '', 3000);
    console.error(error);
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
            'solution-path': showingSolution && isPathCell(rowIndex, colIndex),
            'best-path': showingSolution && isBestPathCell(rowIndex, colIndex)
          }"
          @click="selectCell(rowIndex, colIndex)" 
        >
          <span v-if="cell === 1">I</span>
          <span v-else-if="cell === 2">F</span>
        </div>
      </div>
    </div>
    
    <div v-if="showingSolution && allPaths.length" class="path-navigation">
      <button @click="showPrevPath" :disabled="currentPathIndex === 0" class="nav-button">
        ← Anterior
      </button>
      <span class="path-counter">
        Camino {{ currentPathIndex + 1 }} de {{ allPaths.length }}
      </span>
      <button @click="showNextPath" :disabled="currentPathIndex === allPaths.length - 1" class="nav-button">
        Siguiente →
      </button>
      <button @click="showBestPath" class="best-path-button">
        Mostrar Mejor Camino
      </button>
    </div>
    
    <div class="controls">
      <button v-if="mode && !hasStartPoint" @click="enableStartPointSelection" class="start-button">
        Establecer punto de inicio
      </button> 
      
      <button v-if="mode && hasStartPoint" @click="deleteStartPoint" class="delete-button">
        Borrar punto de inicio
      </button>
      
      <button @click="showSolution" class="solution-button">
        {{ showingSolution ? 'Ocultar Solución' : 'Ver Solución' }}
      </button>

      <button v-if="!IsSaved" @click="saveMaze" class="save-button">
        Guardar Laberinto
      </button>

      <button @click="goHome" class="back-button">
        Volver al Menú
      </button>
    </div>
    
    <div v-if="errorMessage" class="message error-message">
      {{ errorMessage }}
    </div>

    <div v-if="saveMessage" class="message save-message">
      {{ saveMessage }}
    </div>
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
  background-color: #2c3e50;
}

.path {
  background-color: #f8f9fa;
}

.start {
  background-color: #42b883;
  color: white;
}

.end {
  background-color: #ff6b6b;
  color: white;
}

.solution-path {
  background-color: rgba(52, 152, 219, 0.5);
}

.best-path {
  background-color: rgba(46, 204, 113, 0.7);
}

.controls {
  display: flex;
  justify-content: center;
  flex-wrap: wrap;
  gap: 1rem;
  margin: 1.5rem 0;
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

.nav-button {
  padding: 0.6rem 1rem;
  background-color: #3498db;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  transition: background-color 0.3s;
}

.nav-button:hover:not(:disabled) {
  background-color: #3498db;
}

.nav-button:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

.best-path-button {
  padding: 0.6rem 1rem;
  background-color: #2ecc71;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  transition: background-color 0.3s;
}

.best-path-button:hover {
  background-color: #27ae60;
}

button {
  padding: 0.8rem 1.5rem;
  border: none;
  border-radius: 6px;
  font-size: 1rem;
  cursor: pointer;
  transition: all 0.3s ease;
}

.start-button {
  background-color: #42b883;
  color: white;
}

.start-button:hover {
  background-color: #3aa876;
}

.delete-button {
  background-color: #e74c3c;
  color: white;
}

.delete-button:hover {
  background-color: #c0392b;
}

.solution-button {
  background-color: #3498db;
  color: white;
}

.solution-button:hover {
  background-color: #2980b9;
}

.back-button {
  background-color: #f0f0f0;
  color: #2c3e50;
}

.back-button:hover {
  background-color: #e0e0e0;
}

.save-button {
  background-color: #f1c40f;
  color: white;
}

.save-button:hover {
  background-color: #d4ac0d;
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
  background-color: #f1c40f;
  color: white;
}

@media (max-width: 768px) {
  .controls, .path-navigation {
    flex-direction: column;
    align-items: center;
  }
  
  button, .nav-button, .best-path-button {
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