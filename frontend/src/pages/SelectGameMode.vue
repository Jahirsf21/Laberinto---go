<script setup>
import { useRouter } from 'vue-router'


const router = useRouter()

const goHome = () => {
  router.push({ name: 'Home' }) 
}

const goAutoSolve = () => {
    router.push({name: 'AutoSolveDimensions'})
}

const goPlayable = () => {
    router.push({name: 'PlayableDimensions'})
}
</script>

<template>
  <div class="game-mode-container">
    <h1>Selecciona el modo de juego</h1>
    <div class="game-modes">
        <button @click="goAutoSolve" class="mode-button">
        Auto solucionado
      </button>
      <button @click="goPlayable" class="mode-button">
        Modo jugable
      </button>
    </div>
    <button @click="goHome" class="back-button">
      ← Volver a la página principal
    </button>
  </div>
</template>

<style scoped>
.game-mode-container {
  max-width: 600px;
  margin: 0 auto;
  padding: 2rem;
  text-align: center;
}

h1 {
  color: #ffffff;
  margin-bottom: 2rem;
}

.game-modes {
  display: flex;
  flex-direction: column;
  gap: 1rem;
  margin-bottom: 2rem;
}

.mode-button {
  padding: 1rem;
  background-color: #42b883;
  color: white;
  border: none;
  border-radius: 8px;
  font-size: 1.1rem;
  cursor: pointer;
  transition: all 0.3s ease;
}

.mode-button:hover {
  background-color: #3aa876;
  transform: translateY(-2px);
}

.back-button {
  padding: 0.75rem 1.5rem;
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