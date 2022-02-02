package main

import "fmt"

func main() {
	i := 1

	// Go não tem aritmética de ponteiros
	// Para criar um ponteiro, basta usar a notação * ao declarar variável
	var p *int = nil

	// Atribuir endereço da variável i para variável p
	p = &i

	*p++ // Altera valor que está referenciado em p
	i++  // Altera o valor de i

	fmt.Printf(
		"Endereço Ponteiro: %v - Endereço Variável: %v - Valor Ponteiro: %v - Valor Variável: %v",
		i, p, *p, i)

}
