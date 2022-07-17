package main

import (
	"fmt"
	"log"
	"net/http"

	userServer "crud/server"

	"github.com/gorilla/mux"
)

func main() {
	// CRUD

	router := mux.NewRouter()

	router.HandleFunc("/users", userServer.Create).Methods(http.MethodPost)
	router.HandleFunc("/users", userServer.Find).Methods(http.MethodGet)
	router.HandleFunc("/users/{user_id}", userServer.FindById).Methods(http.MethodGet)

	fmt.Println("Listening on port 3000")
	log.Fatal(http.ListenAndServe(":3000", router))

}
