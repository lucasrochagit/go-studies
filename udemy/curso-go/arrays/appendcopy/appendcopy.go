package main

import "fmt"

func main() {
	// append e copy só funcionam com slices
	array1 := [3]int{1, 2, 3}
	var slice1 []int

	// array1 = append(array1, 4, 5, 6) // primeiro argumento do append precisa ser um slice
	slice1 = append(slice1, 5, 6, 7)
	fmt.Println(array1, slice1)

	slice2 := make([]int, 2)
	copy(slice2, slice1)
	fmt.Println(slice2) // slice2 só contem 5 e 6

}
