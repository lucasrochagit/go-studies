package main

import "fmt"

func main() {
	x, y := 1, 2

	// apenas postfix
	x++ // x += 1 ou x = x + 10
	fmt.Println(x)

	y-- // y-=1 ou y = y - 1
	fmt.Println(y)

	// fmt.Println(x == y--) // Não posso fazer incremento/decremento numa expressão
}
