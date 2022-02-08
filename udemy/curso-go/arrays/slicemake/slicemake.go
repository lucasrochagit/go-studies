package main

import "fmt"

func main() {
	// create slice with 10 elem from array with 10 elems
	s := make([]int, 10)

	s[9] = 12
	fmt.Println(s)

	// create slice with 10 elem from array with 20 elems
	s = make([]int, 10, 20)
	fmt.Println(s, len(s), cap(s))

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Println(s, len(s), cap(s))

	s = append(s, 1) // increment array capacity and internal array size
	fmt.Println(s, len(s), cap(s))

}
