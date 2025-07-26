import { defineStore } from 'pinia'
import axios from 'axios'

export const useAuthStore = defineStore('auth', {
  state: () => ({
    token: localStorage.getItem('token') || null,
    role: localStorage.getItem('role') || null,
  }),
  actions: {
    async login(username: string, password: string) {
      const res = await axios.post('/api/login', { username, password })
      this.token = res.data.token
      this.role = res.data.role
      localStorage.setItem('token', this.token)
      localStorage.setItem('role', this.role)
    },

    async register(username: string, password: string) {
      const res = await axios.post('/api/register', { username, password })
      // Optional: auto-login after register
      // await this.login(username, password)
      return res.data // return message/role
    },

    logout() {
      this.token = null
      this.role = null
      localStorage.removeItem('token')
      localStorage.removeItem('role')
    }
  }
})
