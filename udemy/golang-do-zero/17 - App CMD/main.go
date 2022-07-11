package main

import (
	"fmt"
	"linha-de-comando/app"
	"log"
	"os"
)

func main() {
	fmt.Println("Start here...")

	aplicacao := app.Gerar()

	erro := aplicacao.Run(os.Args)
	if erro != nil {
		log.Fatal(erro)
	}
}
