package main

import (
	"fmt"
	"math/rand"
	"time"
)

func escrever(txt string) <-chan string {
	channel := make(chan string)

	go func() {
		for {
			channel <- fmt.Sprintf("Valor recebido: %s", txt)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()
	return channel
}

func multiplexar(chOne, chTwo <-chan string) <-chan string {
	chOut := make(chan string)

	go func() {
		for {
			select {
			case msg := <-chOne:
				chOut <- msg
			case msg := <-chTwo:
				chOut <- msg
			}
		}
	}()

	return chOut
}

func main() {

	channel := multiplexar(escrever("Olá Mundo!"), escrever("Enviando mensagens!"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-channel)
	}

}
