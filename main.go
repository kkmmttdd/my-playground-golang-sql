package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	db, err := sql.Open("mysql",
		"root@tcp(127.0.0.1:3306)/hello")
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		// do something here
		log.Fatal(err)
	}

	defer db.Close()
}
