package router

import (
	"my-blog/handlers"
	"net/http"
	"time"

	gorillahandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
    router := mux.NewRouter()

    // 認証が不要なエンドポイント
    router.HandleFunc("/login", handlers.Login).Methods("POST")
    router.HandleFunc("/logout", handlers.Logout).Methods("POST")
    router.HandleFunc("/check-auth", handlers.CheckAuth).Methods("GET")
    // User routes
    router.HandleFunc("/users", handlers.CreateUser).Methods("POST")
    router.HandleFunc("/users/{id}", handlers.GetUser).Methods("GET")
    // Post routes
    router.HandleFunc("/posts", handlers.GetPosts).Methods("GET")
    router.HandleFunc("/posts/{id}", handlers.GetPost).Methods("GET")
    // Category routes
    router.HandleFunc("/categories", handlers.GetCategories).Methods("GET")
    // Comment routes
    router.HandleFunc("/comments/{post_id}", handlers.GetComments).Methods("GET")

    // 認証が必要なエンドポイント
    authenticatedRouter := router.PathPrefix("/").Subrouter()
    authenticatedRouter.Use(handlers.AuthMiddleware)
    // Post routes
    authenticatedRouter.HandleFunc("/posts", handlers.CreatePost).Methods("POST")
    // r.HandleFunc("/posts/{id}", handlers.UpdatePost).Methods("PUT")
    // r.HandleFunc("/posts/{id}", handlers.DeletePost).Methods("DELETE")
    // Category routes
    authenticatedRouter.HandleFunc("/categories", handlers.CreateCategory).Methods("POST")
    // Comment routes
    authenticatedRouter.HandleFunc("/comments", handlers.CreateComment).Methods("POST")

    return router 
}

func ApplyCORS(router *mux.Router) http.Handler {
    // CORS 設定
    corsOptions := gorillahandlers.CORS(
        gorillahandlers.AllowedOrigins([]string{"http://localhost:3000"}),
        gorillahandlers.AllowedMethods([]string{"OPTIONS", "GET", "POST", "PUT", "DELETE"}),
        gorillahandlers.AllowedHeaders([]string{"Origin", "Content-Type", "Accept"}),
        gorillahandlers.ExposedHeaders([]string{"Content-Length"}),
        gorillahandlers.AllowCredentials(),
        gorillahandlers.MaxAge(int((12 * time.Hour).Seconds())),
    )

    return corsOptions(router)
}
