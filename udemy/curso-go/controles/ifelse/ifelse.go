package main

import "fmt"

func imprimirResultado(nota float64) {
	// if nao tem necessidade de colocar parenteses por padrão
	// não pode existir sem o bloco { }
	if nota >= 7 {
		fmt.Println("Aprovado com nota", nota)
	} else {
		fmt.Println("Reprovado com nota", nota)
	}
}

func main() {
	imprimirResultado(7.3)
	imprimirResultado(5.1)
}
