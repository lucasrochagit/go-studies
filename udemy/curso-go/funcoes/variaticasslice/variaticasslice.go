package main

import "fmt"

func imprimmirAprovados(aprovados ...string) {
	fmt.Println("Lista de aprovados")

	for i, aprovados := range aprovados {
		fmt.Printf("%d) %s\n", i+1, aprovados)
	}
}

func main() {
	aprovados := []string{"José", "Carlos", "João", "Júnior"}
	imprimmirAprovados(aprovados...)
}
