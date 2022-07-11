package main

import (
	"fmt"
	"time"
)

func escrever(txt string) {
	for {
		fmt.Println(txt)
		time.Sleep(time.Second)
	}
}

func main() {
	go escrever("Olá Mundo!")        // go routine
	go escrever("Programando em Go") // go routine
	escrever("Goroutine é mt bom")

}
