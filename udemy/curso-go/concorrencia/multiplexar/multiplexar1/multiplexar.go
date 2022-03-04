package main

import (
	"fmt"
	html "lucasrochagit/go/pkg/mod/github.com/lucasrochagit/go-utils/html"
)

func encaminhar(origem <-chan string, destino chan string) {
	for {
		destino <- <-origem
	}
}

// multiplexar - misturar (mensagens) de canais diferentes em um único canal
func juntar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)
	go encaminhar(entrada1, c)
	go encaminhar(entrada2, c)
	return c
}

func main() {
	t1 := html.Titulo("https://www.cod3r.com.br", "https://www.google.com")
	t2 := html.Titulo("https://www.wikipedia.com", "https://youtube.com")

	c := juntar(t1, t2)

	fmt.Println("1°:", <-c)
	fmt.Println("2°:", <-c)
	fmt.Println("3°:", <-c)
	fmt.Println("4°:", <-c)
}
