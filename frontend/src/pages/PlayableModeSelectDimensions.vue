<script setup>
import { useRouter } from 'vue-router'
import { CreateMaze } from '../../wailsjs/go/main/App'

const router = useRouter()

const dimensions = [5, 10, 15, 20, 25]

const goSelectMode = () => {
  router.push({ name: 'SelectGameMode' })
}

const selectDimension = async (size) => {
  const mode = false; 
  const createdMaze = await CreateMaze(size, mode);
  sessionStorage.setItem('gameData', JSON.stringify({
    maze: createdMaze,
    mode: mode
  }))
  
  router.push({ name: 'Game' });
}
</script>

<template>
  <div class="playable-container">
    <h1>Modo Jugable</h1>
    <p class="subtitle">Elige el tamaño del tablero</p>
    
    <div class="dimensions-grid">
      <button
        v-for="size in dimensions"
        :key="size"
        @click="selectDimension(size)"
        class="dimension-button playable-btn"
      >
        {{ size }}x{{ size }}
      </button>
    </div>
    
    <button @click="goSelectMode" class="back-button">
      ← Volver a selección de modos
    </button>
  </div>
</template>


<style scoped>
.playable-container {
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

.dimensions-grid {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 1rem;
  margin-bottom: 2rem;
}

.dimension-button {
  padding: 1.5rem;
  color: white;
  border: none;
  border-radius: 8px;
  font-size: 1.2rem;
  cursor: pointer;
  transition: all 0.3s ease;
}

.playable-btn {
  background-color: #3498db; 
}

.dimension-button:hover {
  transform: translateY(-3px);
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
}

.playable-btn:hover {
  background-color: #2980b9;
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