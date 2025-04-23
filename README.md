# 📚 Book Store Golang

A simple RESTful API built in Go for managing a collection of books. This project demonstrates basic CRUD operations using Go's standard library, with an in-memory data store for simplicity and educational purposes.

---

## 🚀 Features

- **CRUD Operations**: Create, Read, Update, and Delete books  
- **RESTful API**: Standard HTTP methods for interacting with data  
- **Modular Codebase**: Clean and organized project structure  

---

```yaml
---

## 🛠️ Getting Started

### Prerequisites

- Go 1.16 or higher
```

### Installation & Running

1. Clone the repository and navigate to the project directory:

```bash
git clone https://github.com/Abrahul-107/book_store_Golang.git
cd book_store_Golang
```
Run the application:
```bash
go run cmd/main/main.go
```
The server will start at http://localhost:8080.

🌐 API Endpoints

Method	Endpoint	Description
- GET	/books	Get all books
- GET	/books/{id}	Get a single book by ID
- POST	/books	Create a new book
- PUT	/books/{id}	Update a book by ID
- DELETE	/books/{id}	Delete a book by ID
- 📌 Note: All data is stored in-memory and will be lost upon server restart.

🤝 Contributing
Contributions are welcome! Feel free to fork the repo, open an issue, or submit a pull request.

📄 License
This project is licensed under the MIT License. See the LICENSE file for details.

🔗 Links
GitHub Repository
