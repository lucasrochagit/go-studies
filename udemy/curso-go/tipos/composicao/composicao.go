package main

import "fmt"

type Esportivo interface {
	ligarTurbo()
}

type Luxuoso interface {
	fazerBaliza()
}

type EsportivoLuxoso interface {
	Esportivo
	Luxuoso
	ligarNitro()
}

type BMW7 struct{}

func (b BMW7) ligarTurbo() {
	fmt.Println("Turbo ligado...")
}

func (b BMW7) fazerBaliza() {
	fmt.Println("Fazendo baliza...")
}
func (b BMW7) ligarNitro() {
	fmt.Println("Nitro ligado...")
}

func main() {
	var b EsportivoLuxoso = BMW7{}

	b.ligarTurbo()
	b.fazerBaliza()
	b.ligarNitro()
}
