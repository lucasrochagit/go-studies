package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	var var1 int = 10
	var var2 int = var1

	fmt.Println(var1, "-", var2)

	var1++
	fmt.Println(var1, "-", var2)

	// Ponteiro: Referência de Memória
	var var3 int
	var p1 *int

	fmt.Println(var3, "-", p1) // Resultado: 0 - <nil>

	var3 = 100
	p1 = &var3

	fmt.Println(var3, "-", p1, "-", *p1) // Resultado: 100 - 0x... - 100

	var3 = 150
	fmt.Println(var3, "-", p1, "-", *p1) // Resultado: 150 - 0x... - 150

}
