package main

import (
	"fmt"
	"reflect"
)

func sep() {
	fmt.Println("---------------------------------------")
}

func main() {
	fmt.Println("Arrays e Slices")
	sep()

	// Arrays
	var array1 [5]string
	fmt.Println(array1)
	array1[0] = "1°"
	array1[1] = "2°"
	array1[2] = "3°"
	array1[3] = "4°"
	array1[4] = "5°"
	fmt.Println(array1)
	sep()

	array2 := [5]string{"1°", "2°", "3°", "4°", "5°"}
	fmt.Println(array2)
	sep()

	// é possível definir um array sem definir o tamanho prévio
	// mas o tamanho do array será a quantiade de elementos definidos
	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)
	sep()

	// Slices
	// Slice não é array, embora pareça um array
	//
	slice1 := []int{10, 11, 12, 13, 14, 15, 16, 17}
	fmt.Println(slice1)
	fmt.Println(reflect.TypeOf(array3))
	fmt.Println(reflect.TypeOf(slice1))

	// Para add novos itens num slice, utiliza o Append
	// Será retornado um novo slice, contendo o valor do anterior + valor novo
	// Mas é possível reatribuir o valor do slice para o mesmo slice, se for uma var
	slice2 := append(slice1, 18)
	fmt.Println(slice1, slice2)

	// Slice é uma fatia de um array
	// Nesse caso, pegamos uma fatia de um array existente, o array3
	slice3 := array2[1:3]
	fmt.Println(slice3)
	array2[2] = "Terceiro"
	fmt.Println(slice3)
	sep()

	// Arrays internos
	// make é utilizado para criar um array e retornar um slice.
	slice4 := make([]float32, 10, 11) // cria um array de 11 posições e retorna as 10 primeiras posições
	fmt.Println(
		slice4,
		len(slice4),
		cap(slice4),
	) // len: tamanho e cap: capacidade

	slice4 = append(slice4, 5)
	fmt.Println(
		slice4,
		len(slice4),
		cap(slice4),
	) // len: 11 e cap: 11

	slice4 = append(slice4, 5) // quando estoura a capacidade do slice, o go duplica a capacidade inicial
	fmt.Println(
		slice4,
		len(slice4),
		cap(slice4),
	) // len: 12 e cap: 24

	slice5 := make([]float32, 5) // cria um array de 5 posições e capacidade 5
	fmt.Println(slice5, len(slice5), cap(slice5))

	slice5 = append(slice5, 10)
	fmt.Println(slice5, len(slice5), cap(slice5))
}
