package main

import (
	"log"
	"net/http"
)

func root(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Página Raiz"))
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá mundo!"))
}

func users(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Carregar pagina de usuarios"))
}

func main() {
	http.HandleFunc("/", root)

	http.HandleFunc("/home", home)

	http.HandleFunc("/users", users)

	log.Fatal(http.ListenAndServe(":3000", nil))
}
