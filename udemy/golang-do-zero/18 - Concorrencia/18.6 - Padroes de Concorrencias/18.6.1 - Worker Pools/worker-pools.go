package main

import "fmt"

// funcao que calcula a sequencia de Fibonacci
func fibonacci(pos int) int {
	if pos <= 1 {
		return pos
	}
	return fibonacci(pos-2) + fibonacci(pos-1)
}

// canal recebe dados(tarefas) e envia dados(resultados)
func worker(tarefas <-chan int, resultados chan<- int) {
	for numero := range tarefas {
		resultados <- fibonacci(numero)
	}
}

func main() {

	pos := 40
	tarefas := make(chan int, pos)
	resultados := make(chan int, pos)

	go worker(tarefas, resultados)

	for i := 0; i < pos; i++ {
		tarefas <- i
	}
	close(tarefas)

	for i := 0; i < pos; i++ {
		resultado := <-resultados
		fmt.Printf("%d ", resultado)
	}
	// close(resultados)
}
