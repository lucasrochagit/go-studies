package main

import "fmt"

func main() {
	fmt.Println("Maps")

	// Chave-Valor (tipo JSON)
	// Diferença é que todas as chaves/valores tem que ser do mesmo tipo
	user := map[string]string{
		"first_name": "John",
		"last_name":  "Doe",
	}
	fmt.Println(user)
	fmt.Println(user["first_name"])

	user2 := map[string]map[string]string{
		"name": {
			"first_name": "John",
			"last_name":  "Doe",
		},
		"course": {
			"title":  "Engenharia",
			"campus": "III",
		},
	}

	fmt.Println(user2)
	fmt.Println(user2["name"]["first_name"], "cursa", user2["course"]["title"])

	// para deletar um campo de um map
	delete(user2, "name")
	fmt.Println(user2)

	// para add um campo em um map
	user2["signo"] = map[string]string{
		"nome": "Aquario",
	}
	fmt.Println(user2)
}
