import { createApp } from 'vue'
import App from './App.vue'
import router from './router/index.ts'
import { createPinia } from 'pinia'
import './assets/main.css' // Tailwind import
import './assets/tailwind.css'


const app = createApp(App)

app.use(createPinia())
app.use(router)

app.mount('#app')
