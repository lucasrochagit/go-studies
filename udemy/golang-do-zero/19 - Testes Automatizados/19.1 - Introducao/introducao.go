package main

import (
	"fmt"
	"introducao-testes/enderecos"
)

func main() {
	tipoEndereco := enderecos.TipoEndereco("Praça dos Imigrantes")
	fmt.Println(tipoEndereco)
}
