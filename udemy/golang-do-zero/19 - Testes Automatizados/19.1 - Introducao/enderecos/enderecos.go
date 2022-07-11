package enderecos

import "strings"

// TipoEndereco verifica se o endereço tem um tipo válido e o retorna
func TipoEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}
	primeiraPalavraEndereco := strings.Split(endereco, " ")[0]

	resultado := "Tipo Inválido"

	for _, tipo := range tiposValidos {
		if tipo == strings.ToLower(primeiraPalavraEndereco) {
			resultado = strings.Title(tipo)
		}
	}

	return resultado
}
