package main

import (
    "fmt"
    "os"
    // "log"
    // "reflect"
    "net/http"
    "xyzforum/config"
    "xyzforum/models"
    "xyzforum/handlers"
    "github.com/joho/godotenv"
    // "github.com/google/uuid"
    // _"github.com/go-sql-driver/mysql"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        panic(err)
    }

    config.InitializeTestDB()
    authHandler := handlers.AuthHandler{
        UserModel: &models.UserModel{},
    }

    mux := http.NewServeMux()
    mux.HandleFunc("/register", authHandler.Register)
    mux.HandleFunc("/login", authHandler.Login)

    server := http.Server{
        Addr: os.Getenv("ADDRESS"),
        Handler: mux,
    }

    fmt.Println("Server running on port 5000")
    err = server.ListenAndServe()
    if err != nil {
        panic(err)
    }
}