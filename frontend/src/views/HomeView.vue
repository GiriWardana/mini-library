<template>
  <div class="p-6 max-w-5xl mx-auto">
    <h1 class="text-2xl font-bold mb-4">📚 Library Book List</h1>

    <!-- Book List -->
    <table class="w-full bg-white rounded shadow-md border-collapse mb-10">
      <thead class="bg-gray-100 text-left">
        <tr>
          <th class="p-3">Title</th>
          <th class="p-3">Author</th>
          <th class="p-3">Stock</th>
          <th class="p-3 text-center">Actions</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="book in books" :key="book.id" class="border-t hover:bg-gray-50">
          <td class="p-3">{{ book.title }}</td>
          <td class="p-3">{{ book.author }}</td>
          <td class="p-3">{{ book.stock }}</td>
          <td class="p-3 text-center">
            <button
              v-if="isUser && book.stock > 0"
              @click="borrowBook(book.id)"
              class="bg-green-600 text-white px-3 py-1 rounded hover:bg-green-700"
              :disabled="loadingIds.includes(book.id)"
            >
              {{ loadingIds.includes(book.id) ? 'Processing...' : 'Borrow' }}
            </button>
            <span v-else-if="isUser && book.stock === 0" class="text-sm text-gray-500">Out of Stock</span>
            <span v-else class="text-sm text-gray-400 italic">Login as user to borrow</span>
          </td>
        </tr>
        <tr v-if="books.length === 0">
          <td colspan="4" class="p-4 text-center text-gray-500">No books found.</td>
        </tr>
      </tbody>
    </table>

    <!-- Borrowing History -->
    <div v-if="isUser">
      <h2 class="text-xl font-semibold mb-4">📖 Your Borrowing History</h2>
      <table class="w-full bg-white rounded shadow-md border-collapse">
        <thead class="bg-gray-100 text-left">
          <tr>
            <th class="p-3">Title</th>
            <th class="p-3">Borrowed At</th>
            <th class="p-3">Returned</th>
            <th class="p-3 text-center">Actions</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="borrow in history" :key="borrow.id" class="border-t hover:bg-gray-50">
            <td class="p-3">{{ borrow.book.Title }}</td>
            <td class="p-3">{{ formatDate(borrow.created_at) }}</td>
            <td class="p-3">{{ borrow.returned ? 'Yes' : 'No' }}</td>
            <td class="p-3 text-center">
              <button
                v-if="!borrow.returned"
                @click="returnBook(borrow.id)"
                class="bg-blue-600 text-white px-3 py-1 rounded hover:bg-blue-700"
                :disabled="loadingReturnIds.includes(borrow.id)"
              >
                {{ loadingReturnIds.includes(borrow.id) ? 'Returning...' : 'Return' }}
              </button>
              <span v-else class="text-sm text-gray-400 italic">Returned</span>
            </td>
          </tr>
          <tr v-if="history.length === 0">
            <td colspan="4" class="p-4 text-center text-gray-500">No borrow history.</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import axios from 'axios'
import { useAuthStore } from '../store/auth'

const books = ref([])
const history = ref([])
const loadingIds = ref([])
const loadingReturnIds = ref([])

const auth = useAuthStore()
axios.defaults.headers.common['Authorization'] = `Bearer ${auth.token}`
const isUser = auth.role === 'user'

const fetchBooks = async () => {
  const res = await axios.get('/api/books')
  books.value = res.data.map(book => ({
    id: book.ID,
    title: book.Title,
    author: book.Author,
    stock: book.Stock,
  }))
}

const fetchHistory = async () => {
  if (!isUser) return
  const res = await axios.get('/api/borrow/history', {
    headers: {
      Authorization: `Bearer ${auth.token}`,
    }
  })
  history.value = res.data
}

const borrowBook = async (bookId) => {
  try {
    loadingIds.value.push(bookId)
    const response = await axios.post(`/api/borrow/${bookId}`)
    alert(response.data.message || 'Book borrowed successfully!')
    await fetchBooks()
    await fetchHistory()
  } catch (err) {
    const msg = err.response?.data?.error || err.response?.data?.message || 'Borrow failed'
    alert(msg)
  } finally {
    loadingIds.value = loadingIds.value.filter(id => id !== bookId)
  }
}

const returnBook = async (borrowId) => {
  try {
    loadingReturnIds.value.push(borrowId)
    const res = await axios.post(`/api/borrow/return/${borrowId}`)
    alert(res.data.message || 'Book returned successfully!')
    await fetchBooks()
    await fetchHistory()
  } catch (err) {
    alert(err.response?.data?.error || 'Return failed')
  } finally {
    loadingReturnIds.value = loadingReturnIds.value.filter(id => id !== borrowId)
  }
}

const formatDate = (dateStr) => {
  const date = new Date(dateStr)
  return date.toLocaleString()
}

onMounted(async () => {
  await fetchBooks()
  await fetchHistory()
})
</script>
