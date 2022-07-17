package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // driver de conexão dom o mysql
)

// Connect abre a conexão com o banco de dados
func Connect() (*sql.DB, error) {
	connUrl := "root:root@/devbook?charset=utf8&parseTime=true&loc=Local"
	db, err := sql.Open("mysql", connUrl)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
