-- +goose Up
INSERT INTO books (title, author) VALUES 
('The Go Programming Language', 'Alan Donovan'),
('Go in Action', 'William Kennedy'),
('Introducing Go', 'Caleb Doxsey');

-- +goose Down
DELETE FROM books;
