package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Go MySQL Tutorial")

	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called test
	db, err := sql.Open("mysql", "arunkumar:Tarus_891@tcp(127.0.0.1:3306)/test")

	// if there is an error opening the connection, handle it
	if err != nil {
		log.Println(err.Error())
	}

	// defer the close till after the main function has finished
	// executing
	defer db.Close()

	// perform a db.Query insert
	insert1, err1 := db.Query("INSERT INTO users VALUES ( 'arunkumar' )")
	insert2, err2 := db.Query("INSERT INTO users VALUES ( 'amritha' )")
	// if there is an error inserting, handle it
	if err1 != nil {
		log.Println(err1.Error())
	}
	if err2 != nil {
		log.Println(err2.Error())
	}
	// be careful deferring Queries if you are using transactions
	defer insert1.Close()
	defer insert2.Close()

}
