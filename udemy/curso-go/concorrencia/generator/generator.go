package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

// Google I/0 2012 - Go Concurrency Patterns

//  <-chan - canal somente-leitura

func titulo(urls ...string) <-chan string {
	c := make(chan string)
	for _, url := range urls {
		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)

			r, _ := regexp.Compile("<title>(.*?)<\\/title>")
			c <- r.FindStringSubmatch(string(html))[1]
			// title := r.FindStringSubmatch(string(html))

			// // teste para evitar erro
			// if cap(title) == 0 {
			// 	c <- "Erro ao ler página " + url
			// 	return
			// }

			// c <- title[1]
		}(url)
	}
	return c
}

func main() {
	t1 := titulo("https://www.cod3r.com.br", "https://www.google.com")
	t2 := titulo("https://www.wikipedia.com", "https://youtube.com")

	fmt.Println("Primeiros ", <-t1, "|", <-t2)
	fmt.Println("Segundos: ", <-t1, "|", <-t2)
}
