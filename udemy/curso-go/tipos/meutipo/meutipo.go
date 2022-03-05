package main

import "fmt"

type nota float64

func (n nota) entre(inicio, fim float64) bool {
	return float64(n) >= inicio && float64(n) <= fim
}

func notaParaConceito(n nota) string {
	switch {
	case n.entre(9, 10):
		return "A"
	case n.entre(7, 8.99):
		return "B"
	case n.entre(5, 7.99):
		return "C"
	case n.entre(3, 4.99):
		return "D"
	case n.entre(0, 2.99):
		return "D"
	default:
		return "Nota inválida"
	}
}

func main() {
	var n1, n2, n3 nota = 10, 6, 4
	fmt.Printf("Nota %.1f conceito: %s\n", n1, notaParaConceito(n1))
	fmt.Printf("Nota %.1f conceito: %s\n", n2, notaParaConceito(n2))
	fmt.Printf("Nota %.1f conceito: %s\n", n3, notaParaConceito(n3))
}
