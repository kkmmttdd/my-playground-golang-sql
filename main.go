package main

import (
	"database/sql"
	"fmt"
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
		//log.Println(host)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	// prepared query
	stmt, err := db.Prepare("select host from mysql.user where host like ?")
	if err != nil {
		log.Fatal(err)
	}
	preparedRows, err := stmt.Query("localhost")
	if err != nil {
		log.Fatal(err)
	}
	for preparedRows.Next() {
		preparedRows.Scan(&host)
		fmt.Println(host)
	}
	if err != nil {
		// do something here
		log.Fatal(err)
	}

	// insert statement
	insertStmt, err := db.Prepare("INSERT INTO hello.samples (count, name) VALUES (?, ?)")
	res, err := insertStmt.Exec(2, "name2")
	fmt.Println(res.LastInsertId())
	defer db.Close()
}
