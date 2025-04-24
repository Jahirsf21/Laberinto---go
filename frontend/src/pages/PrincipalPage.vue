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
    const existSavedGames = await MazeExist();
    if (existSavedGames) {
      router.push('SelectGamesSaved');
    } else {
        noSavedGamesMessage.value = 'No hay partidas guardadas';
        setTimeout(() => noSavedGamesMessage.value = '', 3000);
    }
  } catch (error) {
    noSavedGamesMessage.value = 'Error al verificar partidas: ' + error.message;
    setTimeout(() => noSavedGamesMessage.value = '', 3000);
  }
}
</script>

<template>
  <div class="principal-page">
    <h1>Bienvenido a Maze</h1>
    <div class="buttons-container">
      <button @click="startNewGame" class="new-game">Nueva partida</button>
      <button @click="loadGame" class="load-game">Cargar partida</button>
    </div>
    
    <div class="message">
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

.new-game,
.load-game {
  padding: 1.5rem;
  background-color: #42b883;
  color: white;
  border: none;
  border-radius: 8px;
  font-size: 1.2rem;
  cursor: pointer;
  transition: all 0.3s ease;
  width: 200px;
  text-align: center;
}

.new-game:hover,
.load-game:hover {
  background-color: #3aa876;
  transform: translateY(-3px);
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
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