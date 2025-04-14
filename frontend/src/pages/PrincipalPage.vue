<script setup>
import { useRouter } from 'vue-router'
import { MazeExist } from '../../wailsjs/go/main/App'
import { ref } from 'vue'

const router = useRouter()
const noSavedGamesMessage = ref('')

const startNewGame = () => {
  router.push({ name: 'SelectGameMode' })
}

const loadGame = async () => {
  try {
    const existSavedGames = await MazeExist()
    if (existSavedGames) {
      router.push('SelectGamesSaved')
    } else {
      noSavedGamesMessage.value = 'No hay partidas guardadas'
      setTimeout(() => noSavedGamesMessage.value = '', 3000) // Limpiar mensaje después de 3 segundos
    }
  } catch (error) {
    noSavedGamesMessage.value = 'Error al cargar partidas guardadas'
    setTimeout(() => noSavedGamesMessage.value = '', 3000)
  }
}
</script>

<template>
  <div class="principal-page">
    <h1>Bienvenido al laberinto</h1>
    <div class="buttons-container">
      <button @click="startNewGame" class="new-game">Nueva partida</button>
      <button @click="loadGame" class="load-game">Cargar partida</button>
    </div>
    
    <div v-if="noSavedGamesMessage" class="message">
      {{ noSavedGamesMessage }}
    </div>
  </div>
</template>

<style scoped>
.principal-page {
  text-align: center;
  padding: 2rem;
}

.buttons-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 10px;
  margin-top: 20px;
  margin-bottom: 20px;
}

.new-game {
  padding: 12px 24px;
  background-color: #42b883;
  color: white;
  border: none;
  border-radius: 6px;
  font-size: 1rem;
  cursor: pointer;
  transition: background-color 0.3s;
  width: 200px;
}

.new-game:hover {
  background-color: #058938;
}

.load-game {
  padding: 12px 24px;
  background-color: #42b883;
  color: white;
  border: none;
  border-radius: 6px;
  font-size: 1rem;
  cursor: pointer;
  transition: background-color 0.3s;
  width: 200px;
}

.load-game:hover {
  background-color: #058938;
}

.message {
  margin-top: 20px;
  color: #ff4757;
  font-weight: bold;
  animation: fadeInOut 3s ease-in-out;
}

@keyframes fadeInOut {
  0% { opacity: 0; }
  20% { opacity: 1; }
  80% { opacity: 1; }
  100% { opacity: 0; }
}
</style>