package main

import "fmt"

func main() {
	fmt.Printf("Mesma ")
	fmt.Printf("linha.")
	fmt.Println(" Nova")
	fmt.Println("Linha")

	x := 3.141516
	// fmt.Println("O valor de x é " + x) // Não funciona
	xs := fmt.Sprint(x)                 // Parse x para string
	fmt.Println("O valor de x é " + xs) // Agora funciona
	fmt.Println("O valor de x é", x)    // Tbm funciona

	fmt.Printf("O valor de x é %.2f", x) // Duas casas decimais
	fmt.Printf("O valor de x é %.3f", x) // Três casas decimais, arredonda o 1 pra 2

	a, b, c, d := 1, 1.999, false, "opa"
	fmt.Printf("\n%d %f %.1f %t, %s", a, b, b, c, d) // Imprime na forma correta.
	fmt.Printf("\n%v %v %v %v", a, b, c, d)          // %v gancho genérico, serve para diversos tipos
}
