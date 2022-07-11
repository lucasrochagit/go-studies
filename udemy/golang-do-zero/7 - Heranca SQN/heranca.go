package main

import "fmt"

type People struct {
	first_name string
	last_name  string
	age        uint
	height     uint
}

type Student struct {
	People
	course     string
	university string
}

func main() {
	fmt.Println("'Herança'")

	p1 := People{"Lucas", "Rocha", 27, 180}
	e1 := Student{p1, "Computação", "UEPB"}

	fmt.Println(p1)
	fmt.Println(e1)

	e2 := Student{People{"João", "Carlos", 22, 176}, "Design", "UFCG"}
	fmt.Println(e2.first_name, e2.last_name)

}
