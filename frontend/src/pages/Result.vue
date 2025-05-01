<script setup>
import { ref, computed } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()

const savedData = JSON.parse(sessionStorage.getItem('Result'))
const maze = ref(savedData?.maze || [])
const contMoves = ref(savedData?.contMoves || 0) 
const contBestMove = ref(savedData?.contBestMove || 0) 
const contWorstMove = ref(savedData?.contWorstMove || 0) 
const paths = ref(savedData?.paths || [])
const selectedPath = ref(0)
const pathPlayer = ref(savedData?.pathPlayer || [])
const worstPath = ref(savedData?.worstPath || [])
const bestPath = ref(savedData?.bestPath || [])
const showingPlayerPath = ref(false)
const showingBestPath = ref(false)
const showingWorstPath = ref(false)
const showingSolution = ref(false)

const performanceMessage = computed(() => {
  if (contMoves.value === contBestMove.value) {
    return "¡Perfecto! Completaste el laberinto con el camino más corto posible."
  } else if (contMoves.value === contWorstMove.value) {
    return "¡Oh no! Tomaste el camino más largo posible."
  } else if (contMoves.value < contBestMove.value * 1.5) {
    return "¡Buen trabajo! Tu camino fue cercano al óptimo."
  } else if (contMoves.value > contWorstMove.value * 0.8) {
    return "Casi tomas el peor camino, pero lo lograste."
  } else {
    return "Desempeño intermedio. Hay margen para mejorar."
  }
})

const showBestPath = () => {
  showingSolution.value = false
  showingWorstPath.value = false
  showingPlayerPath.value = false
  showingBestPath.value = !showingBestPath.value
}

const showWorstPath = () => {
  showingPlayerPath.value = false
  showingBestPath.value = false
  showingSolution.value = false
  showingWorstPath.value = !showingWorstPath.value
}

const showPlayerPath = () => {
  showingBestPath.value = false
  showingSolution.value = false
  showingWorstPath.value = false
  showingPlayerPath.value = !showingPlayerPath.value
}

const toggleShowPaths = () => {
  showingBestPath.value = false
  showingWorstPath.value = false
  showingPlayerPath.value = false
  showingSolution.value = !showingSolution.value
}

const showNextPath = () => {
  if (selectedPath.value < paths.value.length - 1) selectedPath.value++
}

const showPrevPath = () => {
  if (selectedPath.value > 0) selectedPath.value--
}

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

const isPlayerPathCell = (row, col) => {
  return showingPlayerPath.value && pathPlayer.value?.some(coord => coord.row === row && coord.col === col)
}

const goHome = () => {
  router.push({ name: 'Home' })
}
</script>

<template>
  <div class="result-container">
    <h1>Estadísticas obtenidas</h1>
    <div class="stats">
      <p>Movimientos realizados: {{ contMoves }}</p>
      <p>Movimientos del mejor camino: {{ contBestMove }}</p>
      <p>Movimientos del peor camino: {{ contWorstMove }}</p>
    </div>
    <div class="stats">
      <p>{{ performanceMessage }}</p>
    </div>
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
            'player-path': isPlayerPathCell(rowIndex, colIndex),
            'solution-path': isPathCell(rowIndex, colIndex),
            'best-path': isBestPathCell(rowIndex, colIndex),
            'worst-path': isWorstPathCell(rowIndex, colIndex)
          }"
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

    <div class="controls">
      <button @click="showPlayerPath" class="player-path-button">
        {{ showingPlayerPath ? "Ocultar" : "Mostrar camino realizado" }}
      </button>
      <button @click="showBestPath" class="best-path-button">
        {{ showingBestPath ? "Ocultar" : "Mostrar mejor camino" }}
      </button>
      <button v-if="worstPath.length > bestPath.length" @click="showWorstPath" class="worst-path-button">
        {{ showingWorstPath ? "Ocultar" : "Mostrar peor camino" }}
      </button>
      <button @click="toggleShowPaths" class="solution-button">
        {{ showingSolution ? "Ocultar caminos" : "Mostrar caminos" }}
      </button>
      <button @click="goHome" class="back-button">
        Salir de estadísticas
      </button>
    </div>
  </div>
</template>

<style scoped>

.result-container {
  text-align: center;
  padding: 1rem;
}


.stats {
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
  background-color: #e74c3c;
  color: white;
}

.player-path {
  background-color: rgba(255, 136, 0, 0.8);
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


.controls {
  margin-top: 1rem;
  display: flex;
  justify-content: center;
  gap: 1rem;
}

button {
  padding: 0.6rem 1rem;
  border: none;
  border-radius: 6px;
  background-color: #3498db;
  color: white;
  cursor: pointer;
  transition: all 0.3s ease;
}

button:hover {
  opacity: 0.9;
  transform: translateY(-2px);
}

button:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}


.path-navigation {
  margin: 1rem 0;
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 1rem;
}
</style>