package main

import "fmt"

func main() {
	fmt.Println("Esturturas de Controle")

	numero := 10

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor que 15")
	}

	// if init = declara variável e utiliza no escopo
	// essa variável só é utilizada nesse escopo, não pode
	// ser acessada fora
	if numero2 := numero; numero2 > 0 {
		fmt.Println("Maior que 0")
	} else {
		fmt.Println("Menor que 0")
	}
}
