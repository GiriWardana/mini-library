<template>
  <div class="flex items-center justify-center min-h-screen bg-[#f8f4ee] px-4">
    <div class="w-full max-w-sm p-6 bg-white rounded-xl shadow-lg border border-gray-200">
      <h2 class="text-2xl font-serif font-semibold text-[#264653] mb-6 text-center">📚 Welcome Back</h2>

      <form @submit.prevent="handleLogin" class="space-y-4">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">Username</label>
          <input
            v-model="username"
            class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-[#264653]"
            required
          />
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">Password</label>
          <input
            v-model="password"
            type="password"
            class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-[#264653]"
            required
          />
        </div>

        <button
          type="submit"
          class="w-full bg-[#264653] text-white py-2 rounded-lg hover:bg-[#1f3a46] transition-all font-medium"
        >
          Login
        </button>

        <p v-if="error" class="text-red-500 mt-2 text-sm text-center">{{ error }}</p>
      </form>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '../store/auth'

const username = ref('')
const password = ref('')
const error = ref('')
const router = useRouter()
const auth = useAuthStore()

const handleLogin = async () => {
  try {
    await auth.login(username.value, password.value)
    router.push('/')
  } catch (e) {
    error.value = 'Login failed. Please check your credentials.'
  }
}
</script>
