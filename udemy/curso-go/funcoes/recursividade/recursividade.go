package main

import "fmt"

func fatorial(n int) (int, error) {

	if n < 0 {
		return -1, fmt.Errorf("Número inválido %d", n)
	}

	if n == 0 {
		return 1, nil
	}

	fat, _ := fatorial(n - 1)
	return n * fat, nil
}

func main() {
	result, err := fatorial(0)
	fmt.Printf("Resultado: %d, Erro: %s\n", result, err)
}
