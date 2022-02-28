package main

import (
	"fmt"
	"strings"
)

type Pessoa struct {
	nome      string
	sobrenome string
}

func (p Pessoa) getNomeCompleto() string {
	return p.nome + " " + p.sobrenome
}

func (p *Pessoa) setNomeCompleto(nomeCompleto string) {
	partes := strings.Split(nomeCompleto, " ")
	p.nome, p.sobrenome = partes[0], partes[1]
}

func main() {
	pessoa := Pessoa{nome: "João", sobrenome: "Carlos"}

	nomeCompleto := pessoa.getNomeCompleto()

	fmt.Println(nomeCompleto)

	pessoa.setNomeCompleto("Jorge Aragão")

	fmt.Println(pessoa.getNomeCompleto())

}
