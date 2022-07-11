package main

import "fmt"

func closure() func() {
	// texto que vai ser exibido
	texto := "Dentro da função closure"

	return func() {
		fmt.Println(texto)
	}

}

func main() {
	texto := "Dentro da função main"

	fmt.Println("main >>", texto)

	funcaoNova := closure()
	funcaoNova()
}
