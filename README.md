# just run 
docker compose up 

# CRED (this is default CRED, you can register for the user role after)
login: 
(admin)
username = admin
password = password

(user)
username = user1
password = password


# if using 3 parties 
docker-compose up

# or just using sudo if need permission 
sudo docker compose up 

# you can run 
sudo docker compose -f docker-compose-dev.yaml up
# for fast development / hot reload, but it need some settings in local


####  SPECIFICATION
# 📚 Mini Library Management System

A full-stack mini library management system built with **Golang (Gin)** and **Vue.js (Vue 3)**.  
This system allows user registration, login, role-based access control, and book borrowing/returning functionalities.

---

## 🚀 Overview

This project is a mini library system where:
- Admins can manage books and view all borrowing histories.
- Users can browse books and borrow or return them.

The project includes:
- Backend RESTful API using Golang.
- Frontend web application using Vue 3.
- JWT-based authentication and role-based authorization.

---

## 🎯 Main Objectives

1. Build a RESTful API using Golang.
2. Build a Vue.js web frontend that consumes the API.
3. Implement authentication and role-based authorization using JWT.

---

## ✨ Features

### 1. 🔐 Authentication (Required)
- Register and login endpoints.
- Two user roles: `admin`, `user`.
- Role-based access control:
  - **Admin**: Can add, edit, and delete books.
  - **User**: Can only borrow and return books.

### 2. 📖 Book Management
- **Admin can:**
  - Add new books (title, author, stock).
  - Edit book information.
  - Delete books.
- **All users can:**
  - View list of books.
  - View book details.

### 3. 📦 Borrowing System
- **Users can:**
  - Borrow books (if stock is available).
  - Return borrowed books.
  - View their own borrowing history.
- **Admins can:**
  - View all borrowing history.

---

## 🛠️ Technologies Used

### Backend
- Golang
- Gin Web Framework
- RESTful API
- JWT Authentication
- PostgreSQL

### Frontend
- Vue 3 (Composition API)
- Vue Router
- Axios
- TailwindCSS

---
