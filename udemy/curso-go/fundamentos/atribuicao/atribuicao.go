package main

import "fmt"

func main() {
	var b byte = 3
	fmt.Println(b)

	i := 3 // inferencia de tipo
	i += 3 // i = i+3
	fmt.Println("i ->", i)
	i -= 3 // i = i-3
	fmt.Println("i ->", i)
	i *= 3 // i = i*3
	fmt.Println("i ->", i)
	i /= 3 // i = i/3
	fmt.Println("i ->", i)
	i %= 3 // i = i%3
	fmt.Println("i ->", i)

	x, y := 1, 2 // criar variável
	fmt.Println(x, y)

	x, y = y, x // atribuir variável
	fmt.Println(x, y)
}
