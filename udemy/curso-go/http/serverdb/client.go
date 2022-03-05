package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

// User
type User struct {
	ID   int    `json:"id"`
	Nome string `json:"nome"`
}

// UserHandler analisa o request e delega para a função adequada
func UserHandler(w http.ResponseWriter, r *http.Request) {
	sid := strings.TrimPrefix(r.URL.Path, "/users/")
	id, _ := strconv.Atoi(sid)

	switch {
	case r.Method == "GET" && id > 0:
		findById(w, r, id)
	case r.Method == "GET":
		find(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Página não econtrada")
	}
}

func findById(w http.ResponseWriter, r *http.Request, id int) {
	db, err := sql.Open("mysql", "root:root@/cursogodb")
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Houve um erro interno. Por favor tente novamente mais tarde")
	}
	defer db.Close()

	var u User
	db.QueryRow("select id, nome from usuarios where id = ?", id).Scan(&u.ID, &u.Nome)

	json, _ := json.Marshal(u)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(json))
}

func find(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:root@/cursogodb")
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Houve um erro interno. Por favor tente novamente mais tarde")
	}
	defer db.Close()

	rows, _ := db.Query("select id, nome from usuarios limit 10")
	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User

		rows.Scan(&user.ID, &user.Nome)
		users = append(users, user)
	}

	json, _ := json.Marshal(users)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(json))

}
