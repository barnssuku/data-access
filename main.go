package main

import (
	"database/sql"
	"log"
	"fmt"
	"time"
	_ "github.com/go-sql-driver/mysql"
)

// In production declaring a variable like this for database
// access should be avoided. Placing this in a struct is better.
// This is just for practice.
var db *sql.DB

func main() {
	// Capture connection properties. We need to set th e
	// database log in properties for the application to 
	// connect to the database.
	
	// Get a database handle.
		// This bellow is a beter way to connect to the database 
	// as adviced from the github page of the driver 
	var err error
	db, err = sql.Open("mysql", "data-access-app:data-access-123@/recordings")
	if err != nil {
		log.Fatal(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)


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