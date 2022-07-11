// package enderecos_test // posso colocar o nome do pacote como _test no final
package enderecos

import (
	// . "introducao-testes/enderecos" // quando uso o _test, posso importar os metodos como alias
	"testing"
)

/**
go test ./... --coverprofile cover.txt - gera arquivo txt com dados de cobertura
go tool cover --func=cover.txt - lê e interpreta dados de cobertura no terminal
go tool cover --htmlcover.txt - lê e interpreta dados de cobertura em um html
*/

type CenarioTeste struct {
	info     string
	expected string
}

func TestTipoEndereco(t *testing.T) {
	// t.Parallel() // executa testes em paralelo

	cenariosTeste := []CenarioTeste{
		{"Rua ABC", "Rua"},
		{"RUA DE CIMA", "Rua"},
		{"Avenida Brasil", "Avenida"},
		{"AVENIDA REBOUÇAS", "Avenida"},
		{"Rodovia 200", "Rodovia"},
		{"RODOVIA TRANSAMAZÔNICA", "Rodovia"},
		{"Estrada do Meio", "Estrada"},
		{"ESTRADA DA BANDEIRA", "Estrada"},
		{"Praça da Bandeira", "Tipo Inválido"},
	}

	for _, cenario := range cenariosTeste {
		result := TipoEndereco(cenario.info)
		if result != cenario.expected {
			t.Errorf(
				"O tipo recebido (%s) é diferente do esperado (%s)",
				result,
				cenario.expected,
			)
		}
	}

}

func TestTipoEnderecoVazio(t *testing.T) {
	// t.Parallel() // executa testes em paralelo

	cenario := CenarioTeste{"", "Tipo Inválido"}
	result := TipoEndereco(cenario.info)
	if result != cenario.expected {
		t.Errorf(
			"O tipo recebido (%s) é diferente do esperado (%s)",
			result,
			cenario.expected,
		)
	}
}
