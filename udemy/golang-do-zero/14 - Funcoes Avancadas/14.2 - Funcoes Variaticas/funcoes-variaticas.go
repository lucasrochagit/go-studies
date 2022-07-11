package main

import "fmt"

func soma(numbers ...int) (amount int) {
	for _, num := range numbers {
		amount += num
	}

	return
}

// Eh possivel passar um variático junto de outros campos
// porém só é possível um variático por função
// O variático sempre fica no final
func join(txt string, numbers ...int) {
	for _, num := range numbers {
		fmt.Println(txt, num)
	}
}

func main() {
	amount := soma(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println("amount:", amount)

	join("Hello World! >", 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
}
