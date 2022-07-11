package main

import "fmt"

type User struct {
	nome  string
	idade uint8
}

func (u User) save() {
	fmt.Printf("User saved successfully. Name: %s - Age: %d\n", u.nome, u.idade)
}

func (u User) maiorDeIdade() bool {
	return u.idade >= 18
}

func (u *User) fazerAniversario() {
	u.idade++
}

func main() {
	u1 := User{"User 1", 20}

	fmt.Println(u1)
	u1.save()

	u2 := User{"John", 17}
	u2.save()

	u2Maior := u2.maiorDeIdade()
	fmt.Println("u2 eh maior de idade?", u2Maior)

	u2.fazerAniversario()
	fmt.Println(u1)

	u2MaiorNew := u2.maiorDeIdade()
	fmt.Println("u2 eh maior de idade agora?", u2MaiorNew)

}
