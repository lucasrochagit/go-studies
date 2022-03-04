package main

import (
	"fmt"
	html "lucasrochagit/go/pkg/mod/github.com/lucasrochagit/go-utils/html"
	"time"
)

func oMaisRapido(url1, url2, url3 string) string {
	c1, c2, c3 := html.Titulo(url1), html.Titulo(url2), html.Titulo(url3)

	// esturtua de controle específica para concorrência

	select {
	case t1 := <-c1:
		return t1
	case t2 := <-c2:
		return t2
	case t3 := <-c3:
		return t3
	case <-time.After(500 * time.Millisecond):
		return "Todos perderam!"
		// default:
		// 	return "Sem resposta ainda!"
	}
}

func main() {
	campeao := oMaisRapido(
		"https://www.cod3r.com.br",
		"https://www.google.com",
		"https://www.wikipedia.com",
	)
	fmt.Println(campeao)
}
