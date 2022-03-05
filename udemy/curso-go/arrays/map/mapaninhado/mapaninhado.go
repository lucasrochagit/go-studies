package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela": 12314.00,
			"Gustavo":  4894.48,
		},
		"J": {
			"João": 12314.53,
		},
		"P": {
			"Paulo": 13221.55,
		},
	}

	fmt.Println(funcsPorLetra)

	delete(funcsPorLetra, "P")

	for letra, funcs := range funcsPorLetra {
		fmt.Printf("%s - %v\n", letra, funcs)
	}
}
