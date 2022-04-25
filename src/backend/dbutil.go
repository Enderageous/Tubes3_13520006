package main

import (
	"database/sql"
	"fmt"
	"log"
	_ "net/http"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func InitializeDB() {
	// Load environment variables from .env file
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	cfg := mysql.Config{
		User:   os.Getenv("DB_USER"),
		Passwd: os.Getenv("DB_PASSWORD"),
		Net:    os.Getenv("DB_NET"),
		Addr:   os.Getenv("DB_HOST"),
		DBName: os.Getenv("DB_NAME"),

		AllowNativePasswords: true,
	}
	db, err = sql.Open("mysql", cfg.FormatDSN())
	logError(err)
	pingErr := db.Ping()
	logError(pingErr)
	fmt.Println("Connected to database")
}
