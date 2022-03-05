package main

import "fmt"

func main() {
	// arrays são homogêneos (mesmo tipo) e estáticos (fixos)
	var notas [3]float64
	fmt.Println(notas) // [0 0 0,]

	notas[0], notas[1], notas[2] = 7.7, 10.0, 9.8
	fmt.Println(notas) // [7.7 10 9.8]

	total := 0.0
	for i := 0; i < len(notas); i++ {
		total += notas[i]
	}

	media := total / float64(len(notas))
	fmt.Printf("Média: %.1f\n", media)
}
