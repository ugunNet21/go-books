curl -X POST http://localhost:8080/books/create -d '{"title": "Book Title", "author": "Book Author"}' -H "Content-Type: application/json"
curl http://localhost:8080/books
curl -X POST http://localhost:8080/books/update -d '{"id": 1, "title": "Updated Title", "author": "Updated Author"}' -H "Content-Type: application/json"
curl -X POST http://localhost:8080/books/delete -d '1' -H "Content-Type: application/json"

curl -X POST http://localhost:8080/books/create -H "Content-Type: application/json" -d '{"title": "Book Title", "author": "Author Name"}'


create database dan table ;
CREATE DATABASE go_crud_books;

USE go_crud_books;

CREATE TABLE books (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL
);
