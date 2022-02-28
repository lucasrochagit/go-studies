package main

import "fmt"

type Carro struct {
	nome            string
	velocidadeAtual int
}

type Ferrari struct {
	Carro       // campos anonimos
	turboLigado bool
}

func main() {
	f := Ferrari{}
	f.nome = "F40"
	f.velocidadeAtual = 0
	f.turboLigado = true

	fmt.Printf("A ferrari %s está com o turbo ligado? %v\n", f.nome, f.turboLigado)
	fmt.Println(f)
}
