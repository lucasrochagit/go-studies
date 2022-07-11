package main

import (
	"fmt"
	"time"
)

func main() {
	// i := 0
	// for i < 10 {
	// 	fmt.Println(i, "- incrementando (i)")
	// 	i++
	// 	time.Sleep(time.Second)
	// }
	// fmt.Println(i, "- final (i)")

	// for j := 0; j < 10; j++ {
	// 	fmt.Println(j, "- incrementando (j)")
	// 	time.Sleep(time.Second)
	// }

	names := [3]string{"João", "Davi", "Lucas"}

	for index, name := range names {
		fmt.Println(index, "-", name)
	}

	for index, letter := range "PARALELEPÍPEDO" {
		fmt.Println(index, "-", string(letter)) // se colocar só letra, retorna o código na tabela ascii
	}

	// Não se usa maps em struct, só em maps
	user := map[string]string{
		"first_name": "John",
		"last_name":  "Doe",
	}

	for key, value := range user {
		fmt.Println(key, "-", value)
	}

	// loop infinito

	for {
		fmt.Println("Loop infinito")
		time.Sleep(time.Second / 2)
	}

}
