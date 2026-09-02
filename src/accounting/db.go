package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func connectDB() *sql.DB {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		user,
		password,
		host,
		port,
		database,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Database connection error:", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Database unreachable:", err)
	}

	log.Println("Accounting connected to Percona successfully")
	return db
}
