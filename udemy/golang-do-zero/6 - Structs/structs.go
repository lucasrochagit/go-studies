package main

import "fmt"

func sep() {
	fmt.Println("---------------------------------------")
}

type User struct {
	name    string
	age     uint
	address Address
}

type Address struct {
	street string
	number uint
}

func main() {
	fmt.Println("Arquivo Structs")
	sep()

	// declarar um struct e dps setar os campos
	var u User
	fmt.Println(u)
	u.name = "Jose"
	u.age = 30
	u.address = Address{street: "Rua dos Bobos", number: 0}
	fmt.Println(u)
	sep()

	// declarar um struct com todos os campos
	var u2 User = User{"Lucas", 27, Address{"Rua De Cima", 23}}
	fmt.Println(u2)
	fmt.Println(
		u2.name,
		"tem",
		u2.age,
		"anos de idade e mora na",
		u.address.street,
		"numero",
		u.address.number,
	)
	sep()

	// declarar um struct com um único campo
	u3 := User{name: "John"}
	fmt.Println(u3)
	sep()
}
