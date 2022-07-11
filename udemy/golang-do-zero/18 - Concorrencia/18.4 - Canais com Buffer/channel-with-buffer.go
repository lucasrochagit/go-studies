package main

import "fmt"

func main() {
	channel := make(chan string, 2) // canal com capacidade 2
	channel <- "Hello World!"
	channel <- "Hello World!"
	// channel <- "Hello World!" // extrapolou a capacidade máxima do canal

	close(channel)

	for msg := range channel {
		fmt.Println(msg)
	}
}
