package main

import (
	"database/sql"
	"log"
	"os"
	"fmt"
	"github.com/go-sql-driver/mysql"
)

// In production declaring a variable like this for database
// access should be avoided. Placing this in a struct is better.
// This is just for practice.
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


	// Call db.Ping() to confirm that connection to the database works. At 
	// run time, sql.Open() might not immediately connect depending on the
	// driver. I am using Ping() here to confirm that the database/sql 
	// package can connect when it needs to.
	pingErr := db.Ping()
	// Check for errors from Ping() incase it fails.
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	// Print a message if Ping() connected succesfully.
	fmt.Println("Connected!")
}