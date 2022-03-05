package main

import (
	"fmt"
	"math"
	// m "math" - import as
	// _ "math" import de pacote que não vai usar
)

func main() {
	// sintaxe - const/var nome_da_constante tipo_da_constante = valor
	const PI float64 = 3.14

	// tipo é opcional, o Go faz inferência com base no valor
	var raio = 3.2 // tipo (float64)

	// forma reduzida de criar uma var
	area := PI * math.Pow(raio, 2)

	fmt.Printf("Área do círculo: %.2f", area)

	// declarando constantes em blocos
	const (
		a = 1
		b = 2
	)

	// declarando variáveis em blocos
	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	// declarando variáveis inline
	var e, f bool = true, false
	fmt.Println(e, f)

	g, h, i := 2, false, "epa"
	fmt.Println(g, h, i)
}
