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

    // config.InitializeTestDB()
    config.InitializeDB()

    authHandler := handlers.AuthHandler{
        UserModel: &models.UserModel{},
        AuthValidator: &validators.AuthValidator{},
    }

    forumHandler := handlers.ForumHandler{
        ForumModel: &models.ForumModel{},
        ForumValidator: &validators.ForumValidator{},
    }

    reportForumHandler := handlers.ReportForumHandler{
        ReportForumModel: &models.ReportForumModel{},
        ReportForumValidator: &validators.ReportForumValidator{},
    }

    messageHandler := handlers.MessageHandler{
        MessageModel: &models.MessageModel{},
        MessageValidator: &validators.MessageValidator{},
    }

    reportMessageHandler := handlers.ReportMessageHandler{
        ReportMessageModel: &models.ReportMessageModel{},
        ReportMessageValidator: &validators.ReportMessageValidator{},
    }

    router := mux.NewRouter()
    
    router.HandleFunc("/register", authHandler.Register).Methods("POST")
    router.HandleFunc("/login", authHandler.Login).Methods("POST")
    router.HandleFunc("/profile/{username}", middleware.AuthMiddleware(authHandler.Profile)).Methods("GET")

    router.HandleFunc("/forums", forumHandler.Forums).Methods("GET")
    router.HandleFunc("/forums", middleware.AuthMiddleware(forumHandler.Create)).Methods("POST")
    router.HandleFunc("/forums/{id}", middleware.AuthMiddleware(forumHandler.Update)).Methods("PUT")
    router.HandleFunc("/forums/{id}", middleware.AuthMiddleware(forumHandler.Delete)).Methods("DELETE")
    router.HandleFunc("/forums/{id}/report", middleware.AuthMiddleware(reportForumHandler.Create)).Methods("POST")

    router.HandleFunc("/forums/{id_forum}/messages", messageHandler.ByIdForum).Methods("GET")
    router.HandleFunc("/forums/{id_forum}/messages", middleware.AuthMiddleware(messageHandler.Create)).Methods("POST")
    router.HandleFunc("/forums/{id_forum}/messages/{id}", middleware.AuthMiddleware(messageHandler.Update)).Methods("PUT")
    router.HandleFunc("/forums/{id_forum}/messages/{id}", middleware.AuthMiddleware(messageHandler.Delete)).Methods("DELETE")
    router.HandleFunc("/forums/{id_forum}/messages/{id}/report", middleware.AuthMiddleware(reportMessageHandler.Create)).Methods("POST")

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