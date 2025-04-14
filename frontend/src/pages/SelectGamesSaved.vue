<script setup>
import { useRouter } from 'vue-router'

const router = useRouter()

const goHome = () => {
  router.push({ name: 'Home' }) 
}

</script>


<template>
    <h1>Selecciona la partida</h1>
    <button @click="goHome" class="back-button">
      ← Volver a la página principal
    </button>
</template>

<style scoped>
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