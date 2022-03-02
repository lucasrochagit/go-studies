package main

import (
	"fmt"
	"time"
)

// Channel é a forma de comunicação entre goroutines
// Channel é um tipo (chan)

func doisTresQuatroVezes(base int, c chan int) {
	time.Sleep(1 * time.Second)
	c <- 2 * base

	time.Sleep(1 * time.Second)
	c <- 3 * base

	time.Sleep(1 * time.Second)
	c <- 4 * base
}

func main() {
	ch := make(chan int)

	go doisTresQuatroVezes(2, ch)
	fmt.Println("A")

	a, b := <-ch, <-ch
	fmt.Println("B")
	fmt.Println(a, b)

	fmt.Println(<-ch)
	// fmt.Println(<-ch) // Deadlock - Não existem mais dados para serem recebidos no canal
}
