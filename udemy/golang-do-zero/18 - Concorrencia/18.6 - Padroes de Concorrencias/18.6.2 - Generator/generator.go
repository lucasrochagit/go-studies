package main

import (
	"fmt"
	"time"
)

// Padrão Generator: encapsular o canal
func escrever(txt string) <-chan string {
	channel := make(chan string)

	go func() {
		for {
			channel <- fmt.Sprintf("Valor recebido: %s", txt)
		}
	}()
	return channel
}

func main() {
	channel := escrever("Olá Mundo!")

	for i := 0; i < 10; i++ {
		fmt.Println(<-channel)
		time.Sleep(time.Second / 2)
	}
}
