# 📚 Mini Library Management System

A full-stack mini library system built with **Golang (Gin)** and **Vue.js (Vue 3)**.  
It supports user registration, login, role-based access control, and book borrowing/returning functionalities.

---

## 🚀 Quick Start

### 🔧 Run the App

```bash
docker compose up
```

Then open your browser:

```
http://localhost
```

---

### 🔐 Default Credentials

> You can register a new user (as a regular user). Use the following default credentials for testing:

**Admin:**
- Username: `admin`
- Password: `password`

**User:**
- Username: `user1`
- Password: `password`

---

### 🧪 Development Mode (Hot Reload)

For faster development with hot reload (requires local configuration):

```bash
sudo docker compose -f docker-compose-dev.yaml up
```

---

### 🐳 Running with Docker Compose

If `docker compose` is not available, try:

```bash
docker-compose up
```

Or with elevated permissions:

```bash
sudo docker compose up
```

---

## 🎯 Main Objectives

1. Build a RESTful API using **Golang**.
2. Create a **Vue.js** frontend that consumes the API.
3. Implement **JWT**-based authentication with role-based authorization.

---

## ✨ Features

### 🔐 Authentication & Authorization
- Login & register endpoints.
- Two roles: `admin` and `user`.
- Role-based permissions:
  - **Admin**: Manage books and view all borrowing histories.
  - **User**: Borrow, return, and view their own borrowing history.

### 📖 Book Management
- **Admin** can:
  - Add, edit, and delete books.
- **All users** can:
  - View book lists and details.

### 📦 Borrowing System
- **Users** can:
  - Borrow available books.
  - Return previously borrowed books.
  - Track their borrowing history.
- **Admins** can:
  - View all borrowing histories.

---

## 🛠️ Tech Stack

### Backend
- Golang (Gin)
- PostgreSQL
- JWT for auth
- RESTful API

### Frontend
- Vue 3 (Composition API)
- Vue Router
- Axios
- TailwindCSS

---
