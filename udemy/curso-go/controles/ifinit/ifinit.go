package main

import (
	"fmt"
	"math/rand"
	"time"
)

func numeroAleatorio() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(10)
}

func main() {
	if i := numeroAleatorio(); i > 5 { // tbm é suportado no switch
		fmt.Println(i, "Ganhou!")
	} else {
		fmt.Println(i, "Perdeu!")
	}
	// fmt.Println(i) // não vai funcionar, pois so pertence ao bloco if/else, assim como no for
}
