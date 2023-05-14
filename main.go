package main

import (
    "fmt"
    // "log"
    // "reflect"
    "net/http"
    "xyzforum/config"
    "xyzforum/models"
    "xyzforum/handlers"
    // "github.com/google/uuid"
    // _"github.com/go-sql-driver/mysql"
)

func main() {
    config.InitializeTestDB()
    authHandler := handlers.AuthHandler{
        UserModel: &models.UserModel{},
    }
    mux := http.NewServeMux()
    mux.HandleFunc("/register", authHandler.Register)
    mux.HandleFunc("/login", authHandler.Login)

    server := http.Server{
        Addr: "localhost:5000",
        Handler: mux,
    }

    fmt.Println("Server running on port 5000")
    err := server.ListenAndServe()
    if err != nil {
        panic(err)
    }
}