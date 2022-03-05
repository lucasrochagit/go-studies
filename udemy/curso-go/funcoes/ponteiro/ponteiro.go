package main

import "fmt"

func inc1(n int) {
	n++
}

// revisão: um ponteiro é representado por um *
func inc2(n *int) {
	// revisão: * é usado para desreferenciar, ou seja ter acesso ao valor
	// 			que o ponteiro aponta
	*n++
}

func main() {
	n := 1

	inc1(n)
	fmt.Printf("Primeiro incremento: %d\n", n) // n continua com valor 1

	// revisão: & é usado para obter o endereço da variável por referência
	inc2(&n)
	fmt.Printf("Segundo incremento: %d\n", n) // n tem o valor 2, pq o valor foi alterado na ref

}
