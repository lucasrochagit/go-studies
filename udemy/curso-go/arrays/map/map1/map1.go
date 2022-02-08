package main

import "fmt"

func main() {
	// var aprovados map[int]string // maps devem ser incializados

	aprovados := make(map[int]string)

	aprovados[11122233300] = "Maria"
	aprovados[22233344400] = "Pedro"
	aprovados[33344455500] = "Pedro"

	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(">", aprovados[11122233300])
	delete(aprovados, 11122233300)
	fmt.Println(">", aprovados[11122233300])
}
