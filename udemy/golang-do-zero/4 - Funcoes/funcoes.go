package main

import "fmt"

func Soma(num1, num2 int) int {
	return num1 + num2
}

func MathCalculate(n1, n2 int) (int, int, int, int) {
	soma := n1 + n2
	subtracao := n1 - n2
	multiplicacao := n1 * n2
	divisao := n1 / n2
	return soma, subtracao, multiplicacao, divisao
}

func main() {
	soma := Soma(1, 2)
	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println(txt)
		return ">> " + txt
	}

	resultado := f("Texto da Funcao f")
	fmt.Println(resultado)

	// quando não quero o valor de uma propriedade, usa um _ (como no js)
	soma, sub, mult, div := MathCalculate(10, 5)
	println(soma, sub, mult, div)
}
