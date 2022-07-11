package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

func sep() {
	fmt.Println("---------------------------------------")
}

type Cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	sep()
	c := Cachorro{"Rex", "Dalmata", 3}
	fmt.Println(c)

	cJSONBytes, err := json.Marshal(c)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cJSONBytes) // retorna um slice de uint8 (array de bytes)

	cJSON := bytes.NewBuffer(cJSONBytes)
	fmt.Println(cJSON)
	sep()

	c2 := map[string]string{
		"nome":  "Toby",
		"raca":  "Poodle",
		"idade": "4",
	}
	fmt.Println(c2)
	c2JSONbytes, _ := json.Marshal(c2)
	fmt.Println(c2JSONbytes)

	c2JSON := bytes.NewBuffer(c2JSONbytes)
	fmt.Println(c2JSON)
}
