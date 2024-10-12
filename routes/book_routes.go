package routes

import (
    "net/http"
    "go-crud-books/controllers"
)


func RegisterRoutes(mux *http.ServeMux) {
    mux.HandleFunc("/books", controllers.GetBooks)
    mux.HandleFunc("/books/create", controllers.CreateBook)
    mux.HandleFunc("/books/update", controllers.UpdateBook) // PUT request
    mux.HandleFunc("/books/delete", controllers.DeleteBook) // DELETE request
}
