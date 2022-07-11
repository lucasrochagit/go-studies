package main

import (
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
	cJSONString := `{"nome":"Rex","raca":"Dalmata","idade":3}`
	var c Cachorro

	if err := json.Unmarshal([]byte(cJSONString), &c); err != nil {
		log.Fatal(err)
	}
	fmt.Println(c)
	sep()

	c2JSONString := `{"nome":"Toby","raca":"Poodle","idade":"2"}`
	c2 := make(map[string]string)

	if err := json.Unmarshal([]byte(c2JSONString), &c2); err != nil {
		log.Fatal(err)
	}
	fmt.Println(c2)
}
