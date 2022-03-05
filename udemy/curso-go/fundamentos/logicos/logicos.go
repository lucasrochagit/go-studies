package main

import "fmt"

func comprar(trab1, trab2 bool) (bool, bool, bool) {
	comprarTv50 := trab1 && trab2
	comprarTv32 := trab1 != trab2
	comprarSorvete := trab1 || trab2

	return comprarTv50, comprarTv32, comprarSorvete
}

func main() {
	// tv50, tv32, sorvete := comprar(true, true) // tv50 + sorvete
	// tv50, tv32, sorvete := comprar(true, false) // tv32 + sorvete
	// tv50, tv32, sorvete := comprar(false, true) // tv32 + sorvete
	tv50, tv32, sorvete := comprar(false, false) // nada

	fmt.Printf(
		"Tv50: %t, Tv32: %t, Sorvete: %t, Saudável: %t",
		tv50, tv32, sorvete, !sorvete,
	)

}
