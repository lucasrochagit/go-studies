package main

import "fmt"

type Imprimivel interface {
	toString() string
}

type Pessoa struct {
	nome      string
	sobrenome string
}

type Produto struct {
	nome  string
	preco float64
}

// interfaces são implementadas implicitamente
func (p Pessoa) toString() string {
	return p.nome + " " + p.sobrenome
}

func (p Produto) toString() string {
	return fmt.Sprintf("%s - R$ %.2f", p.nome, p.preco)
}

func imprimir(x Imprimivel) {
	fmt.Println(x.toString())
}

func main() {
	var imprimivel Imprimivel = Pessoa{nome: "João", sobrenome: "Carlos"}
	imprimir(imprimivel)

	imprimivel = Produto{nome: "Notebook", preco: 5000}
	fmt.Println(imprimivel.toString())
	imprimir(imprimivel)

	produto := Produto{nome: "Fone de Ouvido", preco: 200}
	imprimir(produto)
}
