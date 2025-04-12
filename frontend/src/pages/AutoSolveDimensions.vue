<script setup>
import { useRouter } from 'vue-router'
import { CreateMatrix } from '../../wailsjs/go/main/App' // Importación correcta
import { ref } from 'vue'

const router = useRouter()

const goSelectMode = () => {
  router.push({ name: 'SelectGameMode' })
}

const dimensions = [5, 10, 15, 20, 25]
const isLoading = ref(false)
const errorMessage = ref('')

const selectDimension = async (size) => {
  try {
    isLoading.value = true
    errorMessage.value = ''
    
    console.log(`Creando matriz ${size}x${size}...`)
    const matrix = await CreateMatrix(size)
    
    console.log('Matriz creada:', matrix)
    // Aquí puedes hacer algo con la matriz devuelta
    // Por ejemplo, guardarla en un store o navegar a otra vista
    
  } catch (error) {
    console.error('Error al crear matriz:', error)
    errorMessage.value = `Error al crear matriz: ${error.message}`
  } finally {
    isLoading.value = false
  }
}
</script>

<template>
  <div class="auto-solve-container">
    <h1>Modo Auto Solución</h1>
    <p class="subtitle">Selecciona el tamaño del tablero</p>
    
    <div class="dimensions-grid">
      <button
        v-for="size in dimensions"
        :key="size"
        @click="selectDimension(size)"
        class="dimension-button"
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
.auto-solve-container {
  max-width: 600px;
  margin: 0 auto;
  padding: 2rem;
  text-align: center;
}

h1 {
  color: #2c3e50;
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