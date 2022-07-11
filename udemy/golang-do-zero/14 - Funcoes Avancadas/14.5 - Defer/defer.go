package main

import "fmt"

func funcao1() {
	fmt.Println("Executando a funcao 1")
}

func funcao2() {
	fmt.Println("Executando a funcao 2")
}

func aprovado(notas ...int) bool {
	defer fmt.Println("Média calculada. O resultado será exibido")
	fmt.Println("Entrando na funcao para verificar se o aluno foi aprovado")

	media := 0

	for i := range notas {
		media += notas[i]
	}

	return media >= 7
}

func main() {
	// defer funcao1() // defer: adiar a execução da função até o último momento possível
	// funcao2()
	fmt.Println("Aluno Aprovado?", aprovado(7, 6, 5, 8))
}
