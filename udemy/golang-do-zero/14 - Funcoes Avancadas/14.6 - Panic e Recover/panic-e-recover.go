package main

import "fmt"

func recuperarExecucao() {
	fmt.Println("Tentativa de recuperar a execucao")

	if r := recover(); r != nil {
		fmt.Println("Execução recuperada com sucesso")
	}
}

func alunoEstaAprovado(n1, n2 float64) bool {
	defer recuperarExecucao() // executada antes do panic

	media := (n1 + n2) / 2

	if n1 < 0 || n2 < 0 {
		panic("Alguma nota está errada")
	}

	return media > 6
}

func main() {
	fmt.Println(alunoEstaAprovado(-1, 7))
	fmt.Println("Pos execucao!")
}
