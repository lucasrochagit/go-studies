package main

import (
	"errors"
	"fmt"
)

func main() {
	// Inteiros
	// Tipos: int, int8, int16, int32, int64
	// Diferença: Quantidade de bits
	// Tipo int = com base na arquitetura do computador
	// uint -> unsygned int, so aceita positivos (tbm tem 8, 16, 32, 64)
	numero := 100000000000000000
	fmt.Println(numero)

	// var numero2 uint32 = -10000
	var numero2 uint32 = 10000
	fmt.Println(numero2)

	// Alias - Apelido
	// rune - apelido para int32
	// byte = int8
	var numero3 rune = 12345
	fmt.Println(numero3)
	var numero4 byte = 123
	fmt.Println(numero4)

	// Números Reais
	// Tipos: float32, float64
	// Apenas float não é tipo
	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 1045454523.455345
	fmt.Println(numeroReal2)

	// Strings
	// Tipo: string, char
	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto 2"
	fmt.Println(str2)

	char1 := 'B'
	fmt.Println(char1) // << Resultado 66 - posição do B na tabela

	// Valor Zero
	var texto string
	fmt.Println(texto) // << Resultado String em branco

	var inteiro16 int16
	fmt.Println(inteiro16) // << Resultado 0

	var real16 float32
	fmt.Println(real16) // << Resultado 0

	var booleano1 bool
	fmt.Println(booleano1) // << Resultado false

	var erro error
	fmt.Println(erro) // << Resultado <nil>

	var erro2 error = errors.New("Houve um erro interno")
	fmt.Println(erro2) // << Resultado não é uma string, é o tipo erro
}
