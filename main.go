package main

import (
    "fmt"
    "os"
    "net/http"
    "xyzforum/config"
    "xyzforum/models"
    "xyzforum/handlers"
    "xyzforum/middleware"
    "xyzforum/validators"
    "github.com/gorilla/mux"
    "github.com/joho/godotenv"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        panic(err)
    }

    config.InitializeTestDB()

    authHandler := handlers.AuthHandler{
        UserModel: &models.UserModel{},
        AuthValidator: &validators.AuthValidator{},
    }

    forumHandler := handlers.ForumHandler{
        ForumModel: &models.ForumModel{},
        ForumValidator: &validators.ForumValidator{},
    }

    messageHandler := handlers.MessageHandler{
        MessageModel: &models.MessageModel{},
        MessageValidator: &validators.MessageValidator{},
    }

    router := mux.NewRouter()
    router.HandleFunc("/register", authHandler.Register)
    router.HandleFunc("/login", authHandler.Login)
    // router.HandleFunc("/profile", authHandler.Profile)

    router.HandleFunc("/forums", forumHandler.Forums).Methods("GET")
    router.HandleFunc("/forums", middleware.AuthMiddleware(forumHandler.Create)).Methods("POST")
    router.HandleFunc("/forums/{id}", middleware.AuthMiddleware(forumHandler.Update)).Methods("PUT")
    router.HandleFunc("/forums/{id}", middleware.AuthMiddleware(forumHandler.Delete)).Methods("DELETE")

    router.HandleFunc("/forums/{id_forum}/messages", messageHandler.ByIdForum).Methods("GET")
    router.HandleFunc("/forums/{id_forum}/messages", middleware.AuthMiddleware(messageHandler.Create)).Methods("POST")
    router.HandleFunc("/forums/{id_forum}/messages/{id}", middleware.AuthMiddleware(messageHandler.Update)).Methods("PUT")
    router.HandleFunc("/forums/{id_forum}/messages/{id}", middleware.AuthMiddleware(messageHandler.Delete)).Methods("DELETE")

    server := http.Server{
        Addr: os.Getenv("ADDRESS"),
        Handler: router,
    }

    fmt.Println("Server running at: ", os.Getenv("ADDRESS"))
    err = server.ListenAndServe()
    if err != nil {
        panic(err)
    }
}