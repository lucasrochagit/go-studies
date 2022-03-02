package main

import "fmt"

func main() {
	// criando canal que trafega inteiros com buffer tamanho 1
	ch := make(chan int, 1)

	// enviando o dado 1 para o canal
	ch <- 1

	// lendo o dado enviado para o canal
	<-ch

	ch <- 2
	fmt.Println(<-ch)
}
