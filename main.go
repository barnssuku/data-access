package main

import (
	"database/sql"
	"os"
	"log"
	"fmt"
	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	// Capture connection properties. We need to set th e
	// database log in properties for the application to 
	// connect to the database.
	cfg := mysql.Config{
		User: os.Getenv("data-access-app"),
		Passwd: os.Getenv("data-access-123"),
		Net: "tcp",
		Addr: "127.0.0.1:3306",
		DBName: "recordings",
	}

	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Connected!")
}