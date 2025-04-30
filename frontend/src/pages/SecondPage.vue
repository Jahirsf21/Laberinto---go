<script setup>
import { useRouter } from 'vue-router'
import { CreateMaze } from '../../wailsjs/go/main/App'
import { ref } from 'vue'
const router = useRouter()

const dimensions = [5, 10, 15, 20, 25]
const modeSelected = ref(true)
const goHome = () => {
  router.push({ name: 'Home' })
}
  
const selectDimension = async (size) => {
  const createdMaze = await CreateMaze(size, modeSelected.value);
  sessionStorage.setItem('gameData', JSON.stringify({
    maze: createdMaze,
    mode: modeSelected.value
  }))
  
  router.push({ name: 'Game' });
}
</script>

<template>
  <div class="second-page">
    <h1>Maze</h1>
    <p class="subtitle">Elige el tamaño del tablero</p>
    <div class="dimensions-maze">
      <button v-for="size in dimensions" :key="size" @click="selectDimension(size)" class="dimension-button playable-btn"> 
        {{ size }}x{{ size }}
      </button>
    </div>
    <div class="mode-section">
      <p class="subtitle">Selecciona el modo</p>
      <div class="mode-toggle-container">
        <span class="mode-label">Jugable</span>
        <label class="switch">
          <input type="checkbox" v-model="modeSelected">
          <span class="slider round"></span>
        </label>
        <span class="mode-label">Auto Solución</span>
      </div>
    </div>
    <button @click="goHome" class="back-button">
      ← Volver al menu
    </button>
  </div>
</template>

<style scoped>

.second-page {
  text-align: center;
  padding: 2rem;
}

h1 {
  color: #ffffff;
  margin-bottom: 1rem;
  font-size: 2rem;
}

.subtitle {
  color: #7f8c8d;
  margin: 0.5rem 0;
  font-size: 1.1rem;
}

.dimensions-maze {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 1rem;
  margin: 1.5rem 0;
  max-width: 400px;
  margin-left: auto;
  margin-right: auto;
}

.dimension-button {
  padding: 1.5rem;
  background-color: #42b883;
  color: white;
  border: none;
  border-radius: 8px;
  font-size: 1.2rem;
  cursor: pointer;
  transition: all 0.3s ease;
}

.dimension-button:hover {
  background-color: #3aa876;
  transform: translateY(-3px);
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
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

.back-button {
  padding: 1rem 1.5rem;
  background-color: #f0f0f0;
  color: #2c3e50;
  border: none;
  border-radius: 8px;
  font-size: 1.1rem;
  cursor: pointer;
  transition: all 0.3s ease;
  margin-top: 1rem;
  display: inline-flex;
  align-items: center;
  gap: 0.5rem;
}

.back-button:hover {
  background-color: #e0e0e0;
  transform: translateY(-2px);
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
</style>