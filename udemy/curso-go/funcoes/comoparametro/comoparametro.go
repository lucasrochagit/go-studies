package main

import "fmt"

func multiplicacao(a, b int) int { return a * b }

func exec(funcao func(int, int) int, a, b int) int {
	return funcao(a, b)
}

func main() {
	res := exec(multiplicacao, 10, 3)
	fmt.Printf("Resultado: %v\n", res)

	res2 := exec(func(a, b int) int { return a + b }, 2, 3)
	fmt.Printf("Resultado: %v\n", res2)
}
