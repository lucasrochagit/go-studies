package main

import "fmt"

func obterValorAprovado(n int) int {
	defer fmt.Println("Fim!")

	if n >= 5000 {
		fmt.Println("Valor alto...")
	}
	fmt.Println("Valor baixo...")
	return n
}

func main() {
	fmt.Println(obterValorAprovado(60000))
}
