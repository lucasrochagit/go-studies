package main

import "fmt"

func main() {
	p1, p2 := Ponto{0.0, 2.0}, Ponto{2.0, 4.0}

	fmt.Println(catetos(p1, p2))
	fmt.Println(Distancia(p1, p2))
}
