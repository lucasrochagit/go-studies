package main

import "fmt"

func sep() {
	fmt.Println("---------------------------------------")
}

func main() {
	// Operadores Aritméticos
	soma := 1 + 2
	subt := 1 - 2
	mult := 10 * 4
	div := 10 / 2
	mod := 10 % 2

	fmt.Println(soma, subt, mult, div, mod)
	sep()

	// Operadores de Atribuição
	var var1 string = "String1"
	var2 := "String2"

	fmt.Println(var1, var2)
	sep()

	// Operadores Relacionais
	fmt.Println(1 > 2)  // false
	fmt.Println(1 < 2)  // true
	fmt.Println(2 >= 2) // true
	fmt.Println(1 == 2) // false
	fmt.Println(2 <= 2) // true
	fmt.Println(1 != 2) // true
	sep()

	// Operadores Lógicos
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && true) // true
	fmt.Println(true || verdadeiro) // true
	fmt.Println(falso || false)     // false
	sep()

	// Operadores Unários
	numero := 10
	numero++
	fmt.Println(numero) //11
	numero += 15
	fmt.Println(numero) // 26
	numero--
	fmt.Println(numero) // 25
	numero -= 20
	fmt.Println(numero) // 5
	numero *= 5
	fmt.Println(numero) // 25
	//...

}
