<template>
  <div class="min-h-screen bg-[#f8f4ee] text-gray-800 font-sans">
    <!-- Header -->
    <header class="bg-[#264653] text-white px-6 py-4 shadow-md flex justify-between items-center">
      <h1 class="text-2xl font-semibold tracking-wide flex items-center gap-2">
        📚 <span>Mini Library Yuhu</span>
      </h1>
      <nav class="flex gap-4 text-sm items-center">
        <RouterLink
          to="/"
          class="transition-colors hover:text-yellow-300 hover:underline underline-offset-4"
        >Home</RouterLink>

        <RouterLink
          v-if="!auth.token"
          to="/login"
          class="transition-colors hover:text-yellow-300 hover:underline underline-offset-4"
        >Login</RouterLink>

        <RouterLink
          v-if="!auth.token"
          to="/register"
          class="transition-colors hover:text-yellow-300 hover:underline underline-offset-4"
        >Register</RouterLink>

        <RouterLink
          v-if="auth.role === 'admin'"
          to="/admin"
          class="transition-colors hover:text-yellow-300 hover:underline underline-offset-4"
        >Admin Panel</RouterLink>

        <button
          v-if="auth.token"
          @click="handleLogout"
          class="bg-red-500 hover:bg-red-600 text-white text-sm px-3 py-1 rounded-lg transition-all shadow-sm hover:shadow-md"
        >
          Logout
        </button>
      </nav>
    </header>

    <!-- Main -->
    <main class="px-6 py-8 max-w-5xl mx-auto">
      <router-view />
    </main>

    <!-- Footer -->
    <footer class="bg-[#e5e5e5] text-center py-4 text-xs text-gray-600">
      &copy; 2025 Mini Library System • Made with 📖 and ☕
    </footer>
  </div>
</template>

<script setup>
import { RouterLink, useRouter } from 'vue-router'
import { useAuthStore } from './store/auth'

const auth = useAuthStore()
const router = useRouter()

const handleLogout = () => {
  auth.logout()
  router.push('/login')
}
</script>
