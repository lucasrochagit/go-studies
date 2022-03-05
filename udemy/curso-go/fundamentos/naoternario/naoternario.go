package main

import "fmt"

// nao existe operador ternario em go
func obterResultado(nota float64) string {
	// return nota >= 6 ? "Aprovado" : "Reprovado"
	if nota >= 6 {
		return "Aprovado"
	}
	return "Reprovado"
}

func main() {
	fmt.Println(obterResultado(6.2)) // Aprovado
	fmt.Println(obterResultado(5.9)) // Reprovado
}
