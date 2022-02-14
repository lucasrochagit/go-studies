package main

import "fmt"

func troca(p1, p2 int) (segundo int, primeiro int) {
	segundo = p2
	primeiro = p1
	return // retorno limpo
}

func main() {
	a, b := troca(1, 2)
	fmt.Println(a, b)

	c, d := troca(9, 0)
	fmt.Println(c, d)
}
