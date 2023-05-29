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
    "xyzforum/utils"
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

    router.Use(corsOptions)
    
    router.HandleFunc("/register", authHandler.Register).Methods("POST", "OPTIONS")
    router.HandleFunc("/login", authHandler.Login).Methods("POST", "OPTIONS")

    router.HandleFunc("/profile/{username}", middleware.AuthMiddleware(authHandler.Profile)).Methods("GET", "OPTIONS")
    router.HandleFunc("/profile/{username}/nextforum/{created}", middleware.AuthMiddleware(authHandler.NextForum)).Methods("GET", "OPTIONS")
    router.HandleFunc("/profile/{username}/nextmsg/{id_msg}", middleware.AuthMiddleware(authHandler.NextMsg)).Methods("GET", "OPTIONS")

    router.HandleFunc("/forums", forumHandler.Forums).Methods("GET", "OPTIONS")
    router.HandleFunc("/forums", middleware.AuthMiddleware(forumHandler.Create)).Methods("POST", "OPTIONS")
    router.HandleFunc("/forums/{id}", forumHandler.ById).Methods("GET", "OPTIONS")
    router.HandleFunc("/forums/{id}", middleware.AuthMiddleware(forumHandler.Update)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/forums/{id}", middleware.AuthMiddleware(forumHandler.Delete)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/forums/search", forumHandler.FindForum).Methods("POST", "OPTIONS")
    router.HandleFunc("/forums/{id}/searchmsg", forumHandler.FindMsg).Methods("POST", "OPTIONS")
    router.HandleFunc("/forums/{id}/report", middleware.AuthMiddleware(reportForumHandler.Create)).Methods("POST", "OPTIONS")

    router.HandleFunc("/forums/{id_forum}/messages", messageHandler.ByIdForum).Methods("GET", "OPTIONS")
    router.HandleFunc("/forums/{id_forum}/messages", middleware.AuthMiddleware(messageHandler.Create)).Methods("POST", "OPTIONS")
    router.HandleFunc("/forums/{id_forum}/messages/{id}", middleware.AuthMiddleware(messageHandler.Update)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/forums/{id_forum}/messages/{id}", middleware.AuthMiddleware(messageHandler.Delete)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/forums/{id_forum}/messages/{id}/report", middleware.AuthMiddleware(reportMessageHandler.Create)).Methods("POST", "OPTIONS")

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

func corsOptions(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if r.Method == http.MethodOptions {
            utils.WriteResponse(w, r, http.StatusOK, "OK", nil)
            return
        }

        next.ServeHTTP(w, r)
    })
}