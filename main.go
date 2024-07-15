package main

import (
	"log"
	"my-blog/database"
	"my-blog/router"
	"net/http"
)

func main() {
    database.InitDB()
    r := router.InitRouter()
	corsRouter := router.ApplyCORS(r)
    log.Fatal(http.ListenAndServe(":8080", corsRouter))
}
