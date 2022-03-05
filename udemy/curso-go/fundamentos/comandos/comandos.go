package main

import (
	"fmt"
)

/*
Comandos em Go - Terminal

$ go get -v "caminho_da_dependência" - instalar uma dependência
$ go version - versão do go
$ godoc -http=:6060 - habilitar documentação go offline
$ go env - variáveis de ambiente
$ go path - path do go
$ go vet "main.go" - verifica arquivo
$ go build "main.go" - compila código
$ go run "main.go" - compila e executa código
*/

func main() {
	// Printf permite adicionar valores no lugar de valores referenciados por %
	// %s - String
	fmt.Printf("Outro programa em %s!", "GO")

}
