<script setup>
import { useRouter } from 'vue-router'

const router = useRouter()

const startPlaying = () => {
  router.push({ name: 'SelectGameMode' })
}

</script>

<template>
  <div class="principal-page">
    <h1>Bienvenido al laberinto</h1>
    <button @click="startPlaying" class="start-button">
      Empezar a jugar
    </button>
  </div>
</template>

<style scoped>
.principal-page {
  text-align: center;
  padding: 2rem;
}

.start-button {
  padding: 12px 24px;
  background-color: #42b883;
  color: white;
  border: none;
  border-radius: 6px;
  font-size: 1rem;
  cursor: pointer;
  margin-top: 20px;
  transition: background-color 0.3s;
}

.start-button:hover {
  background-color: #3aa876;
}
</style>