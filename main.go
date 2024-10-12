package main

import (
    "log"
    "net/http"
    _"go-crud-books/controllers"
    "go-crud-books/routes"
    "go-crud-books/models"
)

func main() {
    // Initialize the database
    err := models.InitDB("yesjitudb:password@tcp(127.0.0.1:3306)/go_crud_books")
    if err != nil {
        log.Fatal(err)
    }
    routes.RegisterRoutes()
    log.Println("Server is running on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}
