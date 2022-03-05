package main

import (
	"encoding/json"
	"fmt"
)

type Produto struct {
	Id    int      `json:"id"`
	Nome  string   `json:"nome"`
	Preco float64  `json:"preco"`
	Tags  []string `json:"tags"`
}

func main() {
	// struct to json
	var p Produto = Produto{1, "Notebook", 5000, []string{"Promoção", "Eletrônico"}}
	pJSON, _ := json.Marshal(p)
	fmt.Println(string(pJSON))

	// json to struct
	var p2 Produto
	p2String := `{"id":1,"nome":"Café","preco":2.5,"tags":["Promoção","Bebida"]}`
	json.Unmarshal([]byte(p2String), &p2)
	fmt.Println(p2.Tags)
}
