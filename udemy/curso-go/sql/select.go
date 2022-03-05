package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Usuario struct {
	id   int
	nome string
}

func main() {
	db, err := sql.Open("mysql", "root:root@/cursogodb")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, _ := db.Query("select id, nome from usuarios limit 3 offset 1")
	defer rows.Close()

	for rows.Next() {
		var u Usuario
		rows.Scan(&u.id, &u.nome)
		fmt.Println("Usuario:", u)
	}
}
