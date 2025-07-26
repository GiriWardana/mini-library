<template>
  <div class="p-6 max-w-5xl mx-auto">
    <h1 class="text-2xl font-bold mb-6">Admin Panel – Manage Books</h1>

    <table class="w-full bg-white rounded shadow border-collapse">
      <thead class="bg-gray-100">
        <tr>
          <th class="p-3 text-left">Title</th>
          <th class="p-3 text-left">Author</th>
          <th class="p-3 text-left">Stock</th>
          <th class="p-3 text-center">Actions</th>
        </tr>
      </thead>
      <tbody>
        <!-- Editable Rows -->
        <tr v-for="book in books" :key="book.id" class="border-t hover:bg-gray-50">
          <td class="p-2">
            <input
              v-model="book.title"
              class="w-full p-1 border rounded text-sm"
              :readonly="!book.editing"
            />
          </td>
          <td class="p-2">
            <input
              v-model="book.author"
              class="w-full p-1 border rounded text-sm"
              :readonly="!book.editing"
            />
          </td>
          <td class="p-2">
            <input
              v-model.number="book.stock"
              type="number"
              min="0"
              class="w-full p-1 border rounded text-sm"
              :readonly="!book.editing"
            />
          </td>
          <td class="p-2 text-center space-x-2">
            <template v-if="book.editing">
              <button @click="saveEdit(book)" class="text-green-600 hover:underline">Save</button>
              <button @click="cancelEdit(book)" class="text-gray-500 hover:underline">Cancel</button>
            </template>
            <template v-else>
              <button @click="enableEdit(book)" class="text-blue-600 hover:underline">Edit</button>
              <button @click="deleteBook(book.id)" class="text-red-600 hover:underline">Delete</button>
            </template>
          </td>
        </tr>

        <!-- New Book Row -->
        <tr class="border-t bg-gray-50">
          <td class="p-2">
            <input v-model="newBook.title" placeholder="New title" class="w-full p-1 border rounded text-sm" />
          </td>
          <td class="p-2">
            <input v-model="newBook.author" placeholder="New author" class="w-full p-1 border rounded text-sm" />
          </td>
          <td class="p-2">
            <input v-model.number="newBook.stock" placeholder="0" type="number" min="0" class="w-full p-1 border rounded text-sm" />
          </td>
          <td class="p-2 text-center">
            <button @click="addBook" class="text-blue-600 hover:underline">Add</button>
          </td>
        </tr>

        <tr v-if="books.length === 0">
          <td colspan="4" class="text-center p-4 text-gray-500">No books found.</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import axios from 'axios'
import { useAuthStore } from '../store/auth'
const auth = useAuthStore()
if (!auth.token || auth.role !== 'admin') {
  throw new Error('Unauthorized access to Admin Panel')
}
axios.defaults.headers.common['Authorization'] = `Bearer ${auth.token}`

const books = ref([])
const newBook = ref({ title: '', author: '', stock: 0 })

const fetchBooks = async () => {
  const res = await axios.get('/api/books')
  books.value = res.data.map(book => ({
    id: book.ID,
    title: book.Title,
    author: book.Author,
    stock: book.Stock,
    editing: false,
    original: {
      id: book.ID,
      title: book.Title,
      author: book.Author,
      stock: book.Stock,
    }
  }))
}


const enableEdit = (book) => {
  book.editing = true
  book.original = { ...book }
}

const cancelEdit = (book) => {
  book.title = book.original.title
  book.author = book.original.author
  book.stock = book.original.stock
  book.editing = false
}

const saveEdit = async (book) => {
  try {
    await axios.put(`/api/books/${book.id}`, {
      title: book.title,
      author: book.author,
      stock: book.stock
    })
    book.editing = false
    book.original = { ...book }
  } catch (err) {
    alert('Failed to update book')
  }
}

const deleteBook = async (id) => {
  if (confirm('Are you sure you want to delete this book?')) {
    await axios.delete(`/api/books/${id}`)
    await fetchBooks()
  }
}

const addBook = async () => {
  if (!newBook.value.title || !newBook.value.author || newBook.value.stock < 0) {
    alert('Please complete all fields.')
    return
  }
  await axios.post('/api/books', newBook.value)
  newBook.value = { title: '', author: '', stock: 0 }
  await fetchBooks()
}

onMounted(fetchBooks)
</script>
