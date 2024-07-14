package router

import (
	"my-blog/handlers"
	"net/http"
	"time"

	gorillahandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
    r := mux.NewRouter()

    // User routes
    r.HandleFunc("/users", handlers.CreateUser).Methods("POST")
    r.HandleFunc("/users/{id}", handlers.GetUser).Methods("GET")

    // Post routes
    r.HandleFunc("/posts", handlers.CreatePost).Methods("POST")
    r.HandleFunc("/posts", handlers.GetPosts).Methods("GET")
    r.HandleFunc("/posts/{id}", handlers.GetPost).Methods("GET")
    // r.HandleFunc("/posts/{id}", handlers.UpdatePost).Methods("PUT")
    // r.HandleFunc("/posts/{id}", handlers.DeletePost).Methods("DELETE")

    // Category routes
    r.HandleFunc("/categories", handlers.CreateCategory).Methods("POST")
    r.HandleFunc("/categories", handlers.GetCategories).Methods("GET")

    // Comment routes
    r.HandleFunc("/comments", handlers.CreateComment).Methods("POST")
    r.HandleFunc("/comments/{post_id}", handlers.GetComments).Methods("GET")

    return r 
}

func ApplyCORS(router *mux.Router) http.Handler {
    // CORS 設定
    corsOptions := gorillahandlers.CORS(
        gorillahandlers.AllowedOrigins([]string{"http://localhost:3000"}),
        gorillahandlers.AllowedMethods([]string{"OPTIONS", "GET", "POST", "PUT", "DELETE"}),
        gorillahandlers.AllowedHeaders([]string{"X-Requested-With", "Origin", "Content-Type", "Accept"}),
        gorillahandlers.ExposedHeaders([]string{"Content-Length"}),
        gorillahandlers.AllowCredentials(),
        gorillahandlers.MaxAge(int((12 * time.Hour).Seconds())),
    )

    return corsOptions(router)
}
