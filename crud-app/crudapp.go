package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type employee struct {
	id   int
	name string
	city string
}

var (
	id   int
	name string
	city string
)

func main() {
	db, err := sql.Open("mysql", "arunkumar:Tarus_891@tcp(127.0.0.1:3306)/basic")
	if err != nil {
		log.Fatal(err)
	}

	// Create object for employees
	emp1 := employee{1, "arunkumar", "salem"}
	emp2 := employee{2, "amritha", "coimbatore"}
	emp3 := employee{3, "golang", "california"}

	insert1, err := db.Exec("INSERT INTO emp (id, name, city) VALUES (?, ?, ?)", emp1.id, emp1.name, emp1.city)
	if err != nil {
		log.Fatal(err)
	}

	id1, _ := insert1.LastInsertId()
	fmt.Println(id1)

	//	Insert new employee

	insert2, err := db.Exec("INSERT INTO emp (id, name, city) VALUES (?, ?, ?)", emp2.id, emp2.name, emp2.city)
	if err != nil {
		log.Fatal(err)
	}
	id2, _ := insert2.LastInsertId()
	fmt.Println(id2)

	// Insert a new employee

	insert3, err := db.Exec("INSERT INTO emp (id, name, city) VALUES (?,?,?)", emp3.id, emp3.name, emp3.city)
	if err != nil {
		log.Fatal(err)
	}
	id3, _ := insert3.LastInsertId()
	fmt.Println(id3)

	// Update a employee

	upd_query := "update emp SET name = ? WHERE id = ?"
	_, err = db.Exec(upd_query, "tarus", 1)
	if err != nil {
		log.Fatal(err)
	}

	// Select a single employee

	query := "SELECT id, name,city FROM emp WHERE id = ?"
	err = db.QueryRow(query, 1).Scan(&id, &name, &city)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(id, name, city)
	}

	//delete a single user

	_, err = db.Exec("DELETE FROM emp WHERE id = ?", 1)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("deleted from table")
	}

}
