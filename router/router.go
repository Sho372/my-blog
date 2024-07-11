package router

import (
	"my-blog/handlers"

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
