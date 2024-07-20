package main

import (
	"log"
	"my-blog/database"
	"my-blog/middlewares"
	"my-blog/router"
	"net/http"
)

func main() {
    database.InitDB()
    r := router.InitRouter()
	wrappedRouter := middlewares.NewLogger(router.ApplyCORS(r))

    log.Fatal(http.ListenAndServe(":8080", wrappedRouter))
}
