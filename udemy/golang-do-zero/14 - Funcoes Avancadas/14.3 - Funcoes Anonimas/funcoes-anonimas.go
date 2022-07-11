package main

import "fmt"

func main() {
	func() {
		fmt.Println("Ola, mundo")
	}()

	func(txt string) {
		fmt.Println("Texto eh >>", txt)
	}("Ola, mundo")

	result := func(txt string) string {
		return fmt.Sprintf("Recebido -> %s", txt)
	}("Any string")

	fmt.Println(result)
}
