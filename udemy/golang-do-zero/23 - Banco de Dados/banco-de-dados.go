package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	connUrl := "root:root@/devbook?charset=utf8&parseTime=true&loc=Local"
	db, err := sql.Open("mysql", connUrl)
	if err != nil {
		log.Fatal("conn err: ", err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal("ping err: ", err)
	}

	fmt.Println("Connection was opened successfully")

	lines, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal("query err: ", err)
	}
	fmt.Println(lines)
}
