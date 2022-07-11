package main

import (
	"fmt"
	"sync"
	"time"
)

func escrever(txt string) {
	for i := 0; i < 5; i++ {
		fmt.Println(txt)
		time.Sleep(time.Second)
	}
}

func main() {
	var waitGroup sync.WaitGroup

	waitGroup.Add(3)

	go func() {
		escrever("Olá Mundo!")
		waitGroup.Done() // -1
	}()
	go func() {
		escrever("Programando em Go")
		waitGroup.Done() // -1
	}()
	go func() {
		escrever("Goroutine é mt bom")
		waitGroup.Done() // -1
	}()

	// sincroniza 2 ou mais go routines
	waitGroup.Wait()

}
