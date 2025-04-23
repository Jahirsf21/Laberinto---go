<script setup>
import { useRouter } from 'vue-router'
import { ref, watch } from 'vue'
import { SetSPoint, DeleteStartPoint, SaveMaze } from '../../wailsjs/go/main/App'

const router = useRouter()
const savedData = JSON.parse(sessionStorage.getItem('gameData') || 'null')

const maze = ref(savedData?.maze || [])
const mode = ref(savedData?.mode || false)
const IsSaved = ref(savedData?.IsSaved || false)
const selectingStartPoint = ref(false)
const hasStartPoint = ref(false) // Inicializar como false
const errorMessage = ref('')
const showingSolution = ref(false)
const saveMessage = ref('') 

const showSolution = () => {
  showingSolution.value = !showingSolution.value
}

// Actualizar hasStartPoint al cargar el laberinto
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
  const updatedMaze = await SetSPoint(maze.value, row, col)
  maze.value = updatedMaze 
  selectingStartPoint.value = false
}

const deleteStartPoint = async () => {
  const updatedMaze = await DeleteStartPoint(maze.value)
  maze.value = updatedMaze 
}

const saveMaze = async () => {
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
    <h1>Laberinto</h1>
    
    <div class="maze-grid">
      <div 
        v-for="(row, rowIndex) in maze" 
        :key="rowIndex" 
        class="maze-row"
      >
        <div 
          v-for="(cell, colIndex) in row" 
          :key="colIndex"
          class="maze-cell"
          :class="{
            'wall': cell === 3,
            'path': cell === 0,
            'start': cell === 1,
            'end': cell === 2
          }"
          @click="selectCell(rowIndex, colIndex)" 
        >
          <span v-if="cell === 1">I</span>
          <span v-else-if="cell === 2">F</span>
        </div>
      </div>
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
    
    <div v-if="errorMessage" class="message">
      {{ errorMessage }}
    </div>

    <div v-if="saveMessage" class="save-message">
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

.controls {
  display: flex;
  justify-content: center;
  flex-wrap: wrap;
  gap: 1rem;
  margin: 1.5rem 0;
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
  background-color: #42b883;
  color: white;
  display: inline-block;
}

.save-message {
  margin-top: 1rem;
  padding: 0.8rem;
  border-radius: 6px;
  background-color: #f1c40f;
  color: white;
  display: inline-block;
}

@media (max-width: 768px) {
  .controls {
    flex-direction: column;
    align-items: center;
  }
  
  button {
    width: 100%;
    max-width: 250px;
  }
  
  .maze-cell {
    width: 25px;
    height: 25px;
    font-size: 0.8rem;
  }
}
</style>