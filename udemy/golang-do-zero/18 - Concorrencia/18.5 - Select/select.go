package main

import (
	"fmt"
	"time"
)

func main() {
	ch1, ch2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Second / 2)
			ch1 <- "Msg from Channel 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			ch2 <- "Msg from Channel 2"
		}
	}()

	for {
		select { // usado apenas para concorrência
		case msgCh1 := <-ch1:
			fmt.Println(msgCh1)
		case msgCh2 := <-ch2:
			fmt.Println(msgCh2)
		}
	}
}
