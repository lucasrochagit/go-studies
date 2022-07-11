package main

import "fmt"

func diaDaSemana(num int) string {
	switch num {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-Feira"
	case 3:
		return "Terça-Feira"
	case 4:
		return "Quarta-Feira"
	case 5:
		return "Quinta-Feira"
	case 6:
		return "Sexta-Feira"
	case 7:
		return "Sábado"
	default:
		return "Digite um dia válido"
	}
}

func diaDaSemana2(num int) string {
	var diaDaSemana string

	switch {
	case num == 1:
		diaDaSemana = "Domingo"
		fallthrough // força para que o valor caia na proxima condicao
	case num == 2:
		diaDaSemana = "Segunda-Feira"
	case num == 3:
		diaDaSemana = "Terça-Feira"
	case num == 4:
		diaDaSemana = "Quarta-Feira"
	case num == 5:
		diaDaSemana = "Quinta-Feira"
	case num == 6:
		diaDaSemana = "Sexta-Feira"
	case num == 7:
		diaDaSemana = "Sábado"
	default:
		diaDaSemana = "Digite um dia válido"
	}

	return diaDaSemana
}

func main() {
	fmt.Println("Switch")

	dia := 1
	diaDaSemana := diaDaSemana(dia)

	fmt.Println(dia, "-", diaDaSemana)

	diaDaSemana2 := diaDaSemana2(dia)
	fmt.Println(dia, "-", diaDaSemana2)
}
