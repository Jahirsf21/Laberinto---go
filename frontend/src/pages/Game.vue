<script setup>
import { useRouter } from 'vue-router'
import { ref, watch, onMounted } from 'vue'
import {
  SetStartPoint,
  DeleteStartPoint,
  SaveMaze,
  GetPaths,
  GetPath,
  GetBestPath,
  GetWorstPath
} from '../../wailsjs/go/main/App'

// Router and session
const router = useRouter()
const savedData = JSON.parse(sessionStorage.getItem('gameData') || 'null')

// Principal state
const maze = ref(savedData?.maze || [])
const mode = ref(savedData?.mode || false)
const isSaved = ref(savedData?.isSaved || false)

const playerPosition = ref({ row: -1, col: -1 })
const canMove = ref(true)
const contMoves = ref(0)
const contBestMove = ref(0)
const contWorstMove = ref(0)

const pathPlayer = ref([])
const paths = ref([])
const pathToSolution = ref([])
const bestPath = ref([])
const worstPath = ref([])
const selectedPath = ref(0)
const visitedCells = ref([])

// Flags
const hasStartPoint = ref(false)
const selectingStartPoint = ref(false)
const showingSolution = ref(false)
const showingSolutionAnimated = ref(false)
const showingBestPath = ref(false)
const showingWorstPath = ref(false)
const wasShow = ref(false)

// Messages
const errorMessage = ref('')
const saveMessage = ref('')

// Life cycle
onMounted(() => {
  if (!mode.value) {
    const startX = maze.value.findIndex(row => row.includes(1))
    const startY = maze.value[startX]?.indexOf(1)
    playerPosition.value = { row: startX, col: startY }
    pathPlayer.value = [{ row: startX, col: startY }]
    loadPaths()
  }
  window.addEventListener('keydown', handleKeyDown)
})

watch(maze, (newMaze) => {
  hasStartPoint.value = newMaze.some(row => row.includes(1))
}, { immediate: true })

// Funciones principales
const loadPaths = async () => {
  const startX = maze.value.findIndex(row => row.includes(1))
  const startY = maze.value[startX]?.indexOf(1)
  worstPath.value = await GetWorstPath(maze.value, startX, startY)
  bestPath.value = await GetBestPath(maze.value, startX, startY)
  paths.value = await GetPaths(maze.value, startX, startY)
  pathToSolution.value = await GetPath(maze.value, startX, startY)
  contBestMove.value = bestPath.value.length - 1
  contWorstMove.value = worstPath.value.length - 1
}

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
  maze.value = await SetStartPoint(maze.value, row, col)
  await loadPaths()
  selectingStartPoint.value = false
}

const deleteStartPoint = async () => {
  maze.value = await DeleteStartPoint(maze.value)
  showingBestPath.value = false
  showingWorstPath.value = false
  showingSolution.value = false
  bestPath.value = []
  worstPath.value = []
  paths.value = []
}

const handleKeyDown = async (event) => {
  if (!mode.value && canMove.value) {
    switch (event.key) {
      case 'w': movePlayer('up'); break
      case 's': movePlayer('down'); break
      case 'a': movePlayer('left'); break
      case 'd': movePlayer('right'); break
    }
  }
}

const movePlayer = (direction) => {
  let { row, col } = playerPosition.value
  let newRow = row
  let newCol = col

  switch (direction) {
    case 'up': newRow--; break
    case 'down': newRow++; break
    case 'left': newCol--; break
    case 'right': newCol++; break
  }

  if (
    newRow >= 0 && newRow < maze.value.length &&
    newCol >= 0 && newCol < maze.value[0].length &&
    maze.value[newRow][newCol] !== 3
  ) {
    const oldPosition = { row, col }
    playerPosition.value = { row: newRow, col: newCol }

    if (newRow !== oldPosition.row || newCol !== oldPosition.col) {
      contMoves.value++
      pathPlayer.value.push({ row: newRow, col: newCol })
    }

    if (maze.value[newRow][newCol] === 2) {
      goResult()
    }
  }
}

// Cell marking
const isPathCell = (row, col) => {
  if (!showingSolution.value) return false
  const path = paths.value[selectedPath.value]
  return path?.some(coord => coord.fila === row && coord.col === col)
}

const isBestPathCell = (row, col) => {
  return showingBestPath.value && bestPath.value?.some(coord => coord.fila === row && coord.col === col)
}

const isWorstPathCell = (row, col) => {
  return showingWorstPath.value && worstPath.value?.some(coord => coord.fila === row && coord.col === col)
}

// Navigation between paths
const showNextPath = () => {
  if (selectedPath.value < paths.value.length - 1) selectedPath.value++
}

const showPrevPath = () => {
  if (selectedPath.value > 0) selectedPath.value--
}

// Show paths
const toggleShowPaths = async () => {
  if (!hasStartPoint.value) {
    errorMessage.value = 'No puedes ver los caminos sin un punto de inicio'
    setTimeout(() => errorMessage.value = '', 3000)
    return
  }
  if (!mode.value) clearAnimation()
  if (selectingStartPoint.value) selectingStartPoint.value = false
  showingBestPath.value = false
  showingWorstPath.value = false
  if (mode.value && !paths.value.length) await loadPaths()
  showingSolution.value = !showingSolution.value
}

const showPathAnimated = () => {
  canMove.value = false
  showingSolution.value = false
  showingBestPath.value = false
  showingSolutionAnimated.value = true
  wasShow.value = true
  animateSolution()
}

// Animation
let animationInterval = null
let currentStep = 0

const animateSolution = () => {
  if (animationInterval) {
    clearInterval(animationInterval)
    currentStep = 0
  }

  if (!pathToSolution.value?.length) {
    errorMessage.value = 'No hay camino para mostrar'
    setTimeout(() => (errorMessage.value = ''), 3000)
    return
  }

  currentStep = 0
  const start = pathToSolution.value[0]
  playerPosition.value = { row: start.fila, col: start.col }
  visitedCells.value.push({ row: start.fila, col: start.col })

  animationInterval = setInterval(() => {
    currentStep++
    if (currentStep < pathToSolution.value.length) {
      const { fila, col } = pathToSolution.value[currentStep]
      playerPosition.value = { row: fila, col: col }
      if (!visitedCells.value.some(cell => cell.row === fila && cell.col === col)) {
        visitedCells.value.push({ row: fila, col: col })
      }
    } else {
      clearInterval(animationInterval)
      animationInterval = null
      visitedCells.value = []
      showBestPath()
      wasShow.value = true
    }
  }, 300)
}

const clearAnimation = () => {
  if (animationInterval) {
    clearInterval(animationInterval)
    animationInterval = null
  }
  showingSolutionAnimated.value = false
  visitedCells.value = []
  canMove.value = true

  const endRow = maze.value.findIndex(row => row.includes(2))
  if (endRow !== -1) {
    const endCol = maze.value[endRow].indexOf(2)
    playerPosition.value = { row: endRow, col: endCol }
  }
}

const showBestPath = () => {
  if (!hasStartPoint.value) {
    errorMessage.value = 'No puedes ver la solución sin un punto de inicio'
    setTimeout(() => errorMessage.value = '', 3000)
    return
  }
  if (!mode.value) clearAnimation()
  showingSolution.value = false
  showingWorstPath.value = false
  showingBestPath.value = !showingBestPath.value
}

const showWorstPath = () => {
  if (!hasStartPoint.value) {
    errorMessage.value = 'No puedes ver la solución sin un punto de inicio'
    setTimeout(() => errorMessage.value = '', 3000)
    return
  }
  if (!mode.value) clearAnimation()
  showingSolution.value = false
  showingBestPath.value = false
  showingWorstPath.value = !showingWorstPath.value
}

// Save
const saveMaze = async () => {
  if (selectingStartPoint.value) selectingStartPoint.value = false
  const mazeName = prompt("Ingrese el nombre para el laberinto:")
  if (!mazeName) {
    saveMessage.value = 'Debe ingresar un nombre para el laberinto'
    setTimeout(() => saveMessage.value = '', 3000)
    return
  }
  try {
    await SaveMaze(maze.value, mazeName, mode.value)
    saveMessage.value = 'Laberinto guardado correctamente'
    isSaved.value = true
  } catch (error) {
    saveMessage.value = 'Error al guardar el laberinto'
    console.error(error)
  } finally {
    setTimeout(() => saveMessage.value = '', 3000)
  }
}

// Navigation
const goHome = () => {
  router.push({ name: 'Home' })
}

const goResult = () => {
  sessionStorage.setItem('Result', JSON.stringify({
    maze: maze.value,
    contMoves: contMoves.value,
    bestPath: bestPath.value,
    worstPath: worstPath.value,
    paths: paths.value,
    contBestMove: contBestMove.value,
    contWorstMove: contWorstMove.value,
    isSaved: isSaved.value,
    pathPlayer: pathPlayer.value
  }))
  router.push({ name: 'Result' })
}
</script>


<template>
  <div class="game-container">
    <h1>Maze</h1>
    <div class="maze-grid">
      <div v-for="(row, rowIndex) in maze" :key="rowIndex" class="maze-row">
        <div
          v-for="(cell, colIndex) in row"
          :key="colIndex"
          class="maze-cell"
          :class="{
            'wall': cell === 3,
            'path': cell === 0,
            'start': cell === 1,
            'end': cell === 2,
            'solution-path': showingSolution && isPathCell(rowIndex, colIndex),
            'best-path': showingBestPath && isBestPathCell(rowIndex, colIndex),
            'worst-path': showingWorstPath && isWorstPathCell(rowIndex, colIndex),
            'visited': visitedCells.some(c => c.row === rowIndex && c.col === colIndex) 
          }"
          @click="selectCell(rowIndex, colIndex)" 
        >
          <span v-if="cell === 1">I</span>
          <span v-else-if="cell === 2">F</span>
          <div v-if="playerPosition.row === rowIndex && playerPosition.col === colIndex" class="player"></div>
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
    <div class="controls-group">
      <div class="controls">
        <button v-if="mode && !hasStartPoint" @click="enableStartPointSelection" class="start-button">
          Establecer punto de inicio
        </button> 
        <button v-if="mode && hasStartPoint" @click="deleteStartPoint" class="delete-button">
          Borrar punto de inicio
        </button>
        <button v-if="!mode && wasShow === false" @click="showPathAnimated" class="path-animation-button"> 
          {{ showingSolutionAnimated ?"Ocultar" :"Mostrar solución animado" }}
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
      <div class="controls">
        <button v-if="!isSaved" @click="saveMaze" class="save-button">
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
  width: 100%;
  height: 100%;
  position: absolute;
}

.visited {
  background-color: rgba(255, 136, 0, 0.866); 
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

.path-animation-button {
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




@keyframes fadeIn {
  from { opacity: 0; }
  to { opacity: 1; }
}
}
</style>
