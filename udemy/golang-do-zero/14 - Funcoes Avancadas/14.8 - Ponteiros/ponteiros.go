package main

import "fmt"

func inverterSinal(num int) int {
	return num * -1
}

func inverterSinalComPonteiro(num *int) {
	*num = *num * -1
}

func main() {
	num := 20
	numInvertido := inverterSinal(num)

	fmt.Println(num, numInvertido)

	novoNumero := 40
	fmt.Println("before ->", novoNumero)
	inverterSinalComPonteiro(&novoNumero)
	fmt.Println("after ->", novoNumero)
}
