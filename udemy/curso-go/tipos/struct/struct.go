package main

import "fmt"

type Produto struct {
	nome     string
	preco    float64
	desconto float64
}

// Método: função com receiver (receptor)
func (p Produto) precoComDesconto() float64 {
	return p.preco * (1 - p.desconto)
}

func main() {
	var produtoTv Produto = Produto{
		nome:     "TV",
		preco:    2000,
		desconto: 0.10,
	}

	fmt.Println(produtoTv, produtoTv.precoComDesconto())

	produtoNotebook := Produto{"Notebook", 5000, 0.05}

	fmt.Println(produtoNotebook, produtoNotebook.precoComDesconto())

}
