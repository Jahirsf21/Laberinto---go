<script setup>
import { ref, onMounted } from 'vue'
import { GetMazes } from '../../wailsjs/go/main/App'
import { useRouter } from 'vue-router'

const router = useRouter()
const mazes = ref([])

const goHome = () => {
  router.push({ name: 'Home' }) 
}

const loadMazes = async () => {
    const response = await GetMazes()
    mazes.value = response.Mazes || []
}

const selectMaze = (maze) => {
  sessionStorage.setItem('gameData', JSON.stringify({
    maze: maze.matrix,
    mode: maze.mode,
    IsSaved: true
  }))
  router.push({ name: 'Game' })
}

onMounted(() => {
  loadMazes()
})
</script>

<template>
  <div class="maze-selection-container">
    <h1>Selecciona la partida</h1>
    <p class="subtitle">Elige un laberinto guardado</p>
    
    <div class="maze-buttons-grid">
      <div 
        v-for="maze in mazes" 
        :key="maze.name" 
        class="maze-button"
        @click="selectMaze(maze)"
      >
        <h3>{{ maze.name }}</h3>
        <p>Modo: {{ maze.mode ? 'Auto solución' : 'Jugable' }}</p>
        <p>Tamaño: {{ maze.matrix.length }}x{{ maze.matrix[0].length }}</p>
      </div>
    </div>
    
    <button @click="goHome" class="back-button">
      ← Volver al inicio
    </button>
  </div>
</template>

<style scoped>
.maze-selection-container {
  max-width: 600px;
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

.maze-buttons-grid {
  display: grid;
  grid-template-columns: repeat(2, 1fr); 
  gap: 1rem;
  margin-bottom: 2rem;
}

.maze-button {
  padding: 1.5rem;
  color: white;
  background-color: #3498db;
  border: none;
  border-radius: 8px;
  font-size: 1.2rem;
  cursor: pointer;
  transition: all 0.3s ease;
}

.maze-button:hover {
  transform: translateY(-3px);
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  background-color: #2980b9; 
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

.back-button {
  padding: 0.8rem 1.5rem;
  background-color: #f0f0f0;
  color: #2c3e50;
  border: none;
  border-radius: 6px;
  font-size: 1rem;
  cursor: pointer;
  transition: all 0.3s ease;
  display: flex;
  align-items: center;
  gap: 0.5rem;
  margin: 0 auto;
}

.back-button:hover {
  background-color: #e0e0e0;
}
</style>