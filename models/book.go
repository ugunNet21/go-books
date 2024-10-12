package models

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    
)

type Book struct {
    ID     int    `json:"id"`
    Title  string `json:"title"`
    Author string `json:"author"`
}

var db *sql.DB

func InitDB(dataSourceName string) error {
    var err error
    db, err = sql.Open("mysql", dataSourceName)
    if err != nil {
        return err
    }
    // Optionally, you can check the connection
    return db.Ping()
}


func GetAllBooks() ([]Book, error) {
    rows, err := db.Query("SELECT id, title, author FROM books")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var books []Book
    for rows.Next() {
        var book Book
        if err := rows.Scan(&book.ID, &book.Title, &book.Author); err != nil {
            return nil, err
        }
        books = append(books, book)
    }
    return books, nil
}

func CreateBook(book Book) error {
    stmt, err := db.Prepare("INSERT INTO books (title, author) VALUES (?, ?)")
    if err != nil {
        return err
    }

    _, err = stmt.Exec(book.Title, book.Author)
    return err
}