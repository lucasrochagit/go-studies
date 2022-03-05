package main

import "fmt"

func notaParaConceito(nota float64) string {
	var notaInt = int(nota)

	switch notaInt {
	case 10:
		fallthrough // executar o case subsequente
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "E"
	default:
		return "Nota inválida"
	}
}

func main() {
	n1, n2, n3, n4, n5 := 9.1, 8.5, 6.0, 4.4, 2.9
	inv1, inv2 := 11.2, -5.3
	fmt.Printf("Nota: %.1f - Conceito: %v\n", n1, notaParaConceito(n1))
	fmt.Printf("Nota: %.1f - Conceito: %v\n", n2, notaParaConceito(n2))
	fmt.Printf("Nota: %.1f - Conceito: %v\n", n3, notaParaConceito(n3))
	fmt.Printf("Nota: %.1f - Conceito: %v\n", n4, notaParaConceito(n4))
	fmt.Printf("Nota: %.1f - Conceito: %v\n", n5, notaParaConceito(n5))
	fmt.Printf("Nota: %.1f - Conceito: %v\n", inv1, notaParaConceito(inv1))
	fmt.Printf("Nota: %.1f - Conceito: %v\n", inv2, notaParaConceito(inv2))
}
