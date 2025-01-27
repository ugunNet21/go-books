package main

import (
    "log"
    "net/http"
    _ "go-crud-books/controllers"
    "go-crud-books/routes"
    "go-crud-books/models"
)

// CORS middleware to allow requests from your frontend
func cors(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*") // Allow all origins
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

        if r.Method == http.MethodOptions {
            w.WriteHeader(http.StatusNoContent)
            return
        }

        next.ServeHTTP(w, r)
    })
}

func main() {
    // Initialize the database
    err := models.InitDB("yesjitudb:password@tcp(127.0.0.1:3306)/go_crud_books")
    if err != nil {
        log.Fatal(err)
    }

    mux := http.NewServeMux()
    routes.RegisterRoutes(mux) // Register routes with the mux

    // Wrap the mux with CORS middleware
    wrappedMux := cors(mux)

    log.Println("Server is running on :8081")
    if err := http.ListenAndServe(":8081", wrappedMux); err != nil {
        log.Fatal(err)
    }
}
