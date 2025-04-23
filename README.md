# 📚 Book Store Golang

A simple RESTful API built in Go for managing a collection of books. This project demonstrates basic CRUD operations using Go's standard library, with an in-memory data store for simplicity and educational purposes.

---

## 🚀 Features

- **CRUD Operations**: Create, Read, Update, and Delete books
- **RESTful API**: Standard HTTP methods for interacting with data
- **Modular Codebase**: Clean and organized project structure

---

## 📁 Project Structure

book_store_Golang/
├── cmd/
│   └── main/
│       └── main.go       # Entry point of the application
├── pkg/
│   ├── handlers/         # HTTP handlers for API endpoints
│   ├── models/           # Data models representing book entities
│   └── repository/       # In-memory data storage and operations
├── go.mod                # Go module definition
├── go.sum                # Go module checksums
└── README.md             # Project documentation

---

## 🛠️ Getting Started

### Prerequisites

- Go 1.16 or higher

### Installation & Running

1. Clone the repository:

```bash
git clone https://github.com/Abrahul-107/book_store_Golang.git
cd book_store_Golang
go run cmd/main/main.go
```


Method | Endpoint | Description
GET | /books | Get all books
GET | /books/{id} | Get a single book by ID
POST | /books | Create a new book
PUT | /books/{id} | Update a book by ID
DELETE | /books/{id} | Delete a book by ID
