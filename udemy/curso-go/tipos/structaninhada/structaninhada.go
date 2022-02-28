package main

import "fmt"

type Item struct {
	produtoId int
	qtde      int
	preco     float64
}

type Pedido struct {
	userId int
	itens  []Item
}

func (p Pedido) valorTotal() float64 {
	total := 0.0
	for _, item := range p.itens {
		total += item.preco * float64(item.qtde)
	}
	return total
}

func main() {
	pedido := Pedido{
		userId: 1,
		itens: []Item{
			{produtoId: 1, qtde: 10, preco: 20},
			{produtoId: 2, qtde: 100, preco: 25.99},
			{produtoId: 3, qtde: 25, preco: 49.99},
		},
	}

	fmt.Printf("Valor total do pedido é R$ %.2f", pedido.valorTotal())

}
