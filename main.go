package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var (
	//id int
	//name string
	host string
)

func main() {
	db, err := sql.Open("mysql",
		"root@tcp(127.0.0.1:3306)/hello")
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	rows, err := db.Query("select host from mysql.user")
	//rows, err := db.Query("select id, name from users where id = ?", 1)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		//err := rows.Scan(&id, &name)
		err := rows.Scan(&host)
		if err != nil {
			log.Fatal(err)
		}
		//log.Println(id, name)
		log.Println(host)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		// do something here
		log.Fatal(err)
	}

	defer db.Close()
}
