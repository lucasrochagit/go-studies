package main

import "fmt"

type Esportivo interface {
	ligarTurbo()
}

type Ferrari struct {
	modelo          string
	turboLigado     bool
	velocidadeAtual int
}

func (f *Ferrari) ligarTurbo() {
	f.turboLigado = true
}

func main() {
	carro1 := Ferrari{"F40", false, 0}
	carro1.ligarTurbo()

	// var carro2 Esportivo = Ferrari{"F40", false, 0} // Não posso alterar valor do objeto em interface
	var carro2 Esportivo = &Ferrari{"F40", false, 0}
	carro2.ligarTurbo()

	fmt.Println(carro1, carro2)
}
