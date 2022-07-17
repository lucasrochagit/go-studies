package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"crud/database"

	"github.com/gorilla/mux"
)

type user struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Create insere um usuário no banco de dados
func Create(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Falha ao ler o corpo da requisição"))
		return
	}

	var user user

	if err = json.Unmarshal(body, &user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Falha ao converter o usuário para struct"))
		return
	}

	db, err := database.Connect()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Erro ao conectar com o banco de dados"))
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO users (name, email) VALUES (?, ?)")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Erro ao criar o statement"))
		return
	}
	defer stmt.Close()

	insert, err := stmt.Exec(user.Name, user.Email)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Erro ao executar o statement"))
		return
	}

	insertedId, err := insert.LastInsertId()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Erro ao obter o id inserido"))
		return
	}

	user.Id = insertedId
	result, _ := json.Marshal(user)

	w.WriteHeader(http.StatusCreated)
	w.Write(result)
}

// Find traz todos os usuários salvos do banco de dados
func Find(w http.ResponseWriter, r *http.Request) {
	db, err := database.Connect()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Erro ao conectar com o banco de dados"))
		return
	}
	defer db.Close()

	lines, err := db.Query("SELECT * FROM users LIMIT 100 OFFSET 0;")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Erro ao buscar os usuários no banco de dados"))
		return
	}
	defer lines.Close()

	var users []user
	for lines.Next() {
		var user user
		if err := lines.Scan(&user.Id, &user.Name, &user.Email); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Erro ao escanear usuários"))
			return
		}
		users = append(users, user)
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(users); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Erro ao converter usuários para JSON"))
		return
	}
	return

}

// FindById traz um usuário salvo no banco de dados através do id
func FindById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userId, err := strconv.ParseUint(params["user_id"], 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Erro ao recuperar o id do usuário"))
		return
	}

	db, err := database.Connect()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Erro ao conectar com o banco de dados"))
		return
	}
	defer db.Close()

	line, err := db.Query("SELECT * FROM users WHERE id=?", userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Erro ao buscar o usuário no banco de dados"))
		return
	}
	defer line.Close()

	var user user
	if line.Next() {
		if err := line.Scan(&user.Id, &user.Name, &user.Email); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Erro ao escanear usuário"))
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Erro ao converter usuário para JSON"))
		return
	}
	return
}
