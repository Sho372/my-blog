package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbName := os.Getenv("DB_NAME")

    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)

    var err error
    DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{DisableAutomaticPing: true})
    if err != nil {
        log.Fatalf("Could not connect to the database: %v", err)
    }

    fmt.Println("Successfully connected to the database!")
}
