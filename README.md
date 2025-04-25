# 📚 Go Book API

A simple RESTful API built with **Go** and **SQLite**, using `net/http` and `goose` for database migrations.

---

## ✨ Features

- CRUD operations for books
- SQLite database with migrations
- RESTful endpoints
- Written in pure Go (no external web frameworks)

---

## 📡 API Endpoints
```
Method	Endpoint	Description
GET	    /books	    Get all books
GET	    /books/{id}	Get book by ID
POST	/books	    Create a new book
```

---

## 🧱 Project Structure
```
go-books/
├── books.db
├── db
│   └── migrations
│       ├── 20250425150826_create_books_table.sql
│       └── 20250425151233_seed_books.sql
├── go.mod
├── go.sum
├── main.go
├── run_migrations.go
└── seed
    └── seed.go
```

---

## 📜 License

MIT

---

## 👋 Contributing

PRs are welcome! Fork it, branch it, hack it, and send it 🚀