package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	// cuidado ...
	fmt.Println("Teste " + string(97))

	// inteiro pra string
	fmt.Println("Teste " + strconv.Itoa(123))

	// string pra int
	num, _ := strconv.Atoi(("123"))
	fmt.Println(num - 122)

	b, _ := strconv.ParseBool("true")
	// b, _ :=strconv.ParseBool("false") // não printa verdaeiro
	// b, _ :=strconv.ParseBool("1") // printa verdadeiro
	// b, _ :=strconv.ParseBool("0") // não printa verdadeiro
	if b {
		fmt.Println("Verdadeiro")
	}
}
