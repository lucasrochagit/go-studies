package main

import "fmt"

// funcao que calcula a sequencia de Fibonacci
func fibonacci(pos int) int {
	if pos <= 1 {
		return pos
	}
	return fibonacci(pos-2) + fibonacci(pos-1)
}

func main() {

	pos := 40
	for i := 0; i < pos; i++ {
		fmt.Printf("%d ", fibonacci(i))
	}
}
