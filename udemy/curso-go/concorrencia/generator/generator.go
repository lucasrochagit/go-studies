package main

import (
	"fmt"
	html "lucasrochagit/go/pkg/mod/github.com/lucasrochagit/go-utils/html"
)

func main() {
	t1 := html.Titulo("https://www.cod3r.com.br", "https://www.google.com")
	t2 := html.Titulo("https://www.wikipedia.com", "https://youtube.com")

	fmt.Println("Primeiros ", <-t1, "|", <-t2)
	fmt.Println("Segundos: ", <-t1, "|", <-t2)
}
