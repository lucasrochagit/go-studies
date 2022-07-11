package main

import (
	"fmt"
	"time"
)

func escrever(txt string, channel chan string) {
	time.Sleep(2 * time.Second)
	for i := 0; i < 5; i++ {
		channel <- txt
		time.Sleep(time.Second)
	}

	close(channel)
}

func main() {
	channel := make(chan string)
	go escrever("Olá mundo", channel)

	fmt.Println("Depois da função escrever começar a ser executada")

	// deadlock! canal aberto sem identificar msgs
	// deadlocs são identificados apenas em execução, não em compilação
	// for {
	// 	msg, isChannelOpen := <-channel
	// 	fmt.Println(msg)
	// }

	// Solução para o Deadlock 01
	// for {
	// 	msg, isChannelOpen := <-channel
	// 	if !isChannelOpen {
	// 		break
	// 	}
	// 	fmt.Println(msg)
	// }

	// Solução 02
	for message := range channel {
		fmt.Println(message)
	}

	fmt.Println("Fim da execução!")
}
