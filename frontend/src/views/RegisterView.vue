<template>
  <div class="flex flex-col items-center justify-center min-h-screen bg-gray-50">
    <div class="w-full max-w-sm p-6 bg-white rounded shadow-md">
      <h2 class="text-xl font-bold mb-4">Register</h2>
      <form @submit.prevent="handleRegister">
        <div class="mb-4">
          <label class="block mb-1 font-medium">Username</label>
          <input v-model="username" class="w-full px-3 py-2 border rounded" required />
        </div>
        <div class="mb-4">
          <label class="block mb-1 font-medium">Password</label>
          <input v-model="password" type="password" class="w-full px-3 py-2 border rounded" required />
        </div>
        <button type="submit" class="w-full bg-green-600 text-white py-2 rounded hover:bg-green-700">
          Register
        </button>
        <p v-if="error" class="text-red-500 mt-2 text-sm">{{ error }}</p>
        <p v-if="success" class="text-green-600 mt-2 text-sm">{{ success }}</p>
      </form>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import axios from 'axios'

const username = ref('')
const password = ref('')
const error = ref('')
const success = ref('')
const router = useRouter()

const handleRegister = async () => {
  error.value = ''
  success.value = ''
  try {
    const res = await axios.post('/api/register', {
      username: username.value,
      password: password.value,
    })
    success.value = 'Register success! Redirecting to login...'
    setTimeout(() => router.push('/login'), 1500)
  } catch (e) {
    if (e.response?.status === 409) {
      error.value = 'Username already exists.'
    } else {
      error.value = 'Registration failed. Please try again.'
    }
  }
}
</script>
