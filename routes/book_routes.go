package routes

import (
    "net/http"

    "go-crud-books/controllers"
)

func RegisterRoutes() {
    http.HandleFunc("/books", controllers.GetBooks)
    http.HandleFunc("/books/create", controllers.CreateBook)
    http.HandleFunc("/books/update", controllers.UpdateBook)
    http.HandleFunc("/books/delete", controllers.DeleteBook)
}
