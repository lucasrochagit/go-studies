package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"Jose":   11321.50,
		"Joao":   3213.34,
		"Carlos": 1200.00,
	}

	funcsESalarios["Rafael"] = 1350.00
	fmt.Println(funcsESalarios)

	delete(funcsESalarios, "Inexistentes")

	for nome, salario := range funcsESalarios {
		fmt.Printf("\n%s - RS %.2f\n", nome, salario)
	}
}
