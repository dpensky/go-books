-- +goose Up
CREATE TABLE books (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  title TEXT NOT NULL,
  author TEXT NOT NULL
);

-- +goose Down
DROP TABLE books;
