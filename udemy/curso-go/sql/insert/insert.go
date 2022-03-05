package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@/cursogodb")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	stmt, _ := db.Prepare("insert into usuarios(nome) values(?)")
	stmt.Exec("Maria")
	stmt.Exec("João")
	res, _ := stmt.Exec("Pedro")

	id, _ := res.LastInsertId()
	fmt.Printf("Último id inserido: %d\n", id)

	linhas, _ := res.RowsAffected()
	fmt.Printf("Quantidade de linhas afetadas no último comando: %d\n", linhas)

}
