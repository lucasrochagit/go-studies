package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

type User struct {
	Nome  string
	Email string
}

func root(w http.ResponseWriter, r *http.Request) {
	if templates != nil {
		templates.ExecuteTemplate(w, "home.html", nil)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	if templates != nil {
		u := User{"John", "john@mail.com"}
		templates.ExecuteTemplate(w, "home.html", u)
	}
}

func main() {
	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/", home)

	fmt.Println("Running on port 3000...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
