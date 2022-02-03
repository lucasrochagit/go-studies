package main

import "fmt"

func notaParaConceito(nota float64) string {

	switch {
	case nota >= 9 && nota <= 10:
		return "A"
	case nota >= 8 && nota < 9:
		return "B"
	case nota >= 5 && nota < 8:
		return "C"
	case nota >= 3 && nota < 5:
		return "D"
	default:
		return "E"
	}
}

func main() {
	n1, n2, n3, n4, n5 := 9.1, 8.5, 6.0, 4.4, 2.9
	fmt.Printf("Nota: %.1f - Conceito: %v\n", n1, notaParaConceito(n1))
	fmt.Printf("Nota: %.1f - Conceito: %v\n", n2, notaParaConceito(n2))
	fmt.Printf("Nota: %.1f - Conceito: %v\n", n3, notaParaConceito(n3))
	fmt.Printf("Nota: %.1f - Conceito: %v\n", n4, notaParaConceito(n4))
	fmt.Printf("Nota: %.1f - Conceito: %v\n", n5, notaParaConceito(n5))
}
