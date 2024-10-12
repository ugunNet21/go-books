package controllers

import (
    "encoding/json"
    "net/http"
    "sync"

    "go-crud-books/models"
)

var (
    books  = make([]models.Book, 0)
    nextID = 1
    mu     sync.Mutex
)

// GetBooks untuk mendapatkan semua buku
func GetBooks(w http.ResponseWriter, r *http.Request) {
    books, err := models.GetAllBooks()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(books)
}

// CreateBook untuk membuat buku baru
func CreateBook(w http.ResponseWriter, r *http.Request) {
    var book models.Book
    if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    mu.Lock()
    book.ID = nextID
    nextID++
    books = append(books, book)
    mu.Unlock()

    if err := models.CreateBook(book); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(book)
}

// UpdateBook untuk memperbarui buku
func UpdateBook(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPut {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    var updatedBook models.Book
    if err := json.NewDecoder(r.Body).Decode(&updatedBook); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if err := models.UpdateBook(updatedBook); err != nil {
        http.Error(w, "Book not found", http.StatusNotFound)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(updatedBook)
}

// DeleteBook untuk menghapus buku
func DeleteBook(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodDelete {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    var book models.Book
    if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if err := models.DeleteBook(book.ID); err != nil {
        http.Error(w, "Book not found", http.StatusNotFound)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}
