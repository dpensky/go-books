package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "books.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(`
        DROP TABLE IF EXISTS books;
        CREATE TABLE books (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            title TEXT NOT NULL,
            author TEXT NOT NULL
        );

        INSERT INTO books (title, author) VALUES 
        ('The Go Programming Language', 'Alan Donovan'),
        ('Go in Action', 'William Kennedy'),
        ('Introducing Go', 'Caleb Doxsey');
    `)
	if err != nil {
		log.Fatal("Failed to seed database:", err)
	}

	log.Println("books.db created and seeded successfully!")
}
