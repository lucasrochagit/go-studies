package main

import "fmt"

func f1() {
	fmt.Println("F1")
}

func f2(p1, p2 string) {
	fmt.Printf("F2: %s %s \n", p1, p2)
}

func f3() string {
	return "F3"
}

func f4(p1, p2 string) string {
	return fmt.Sprintf("F4: %s %s", p1, p2)
}

func f5() (string, string) {
	return "Retorno 1", "Retorno 2"
}

func main() {
	f1()
	f2("p1", "p2")
	fmt.Printf("F3 %s\n", f3())
	fmt.Printf("F4 %s\n", f4("p1", "p2"))

	r1, r2 := f5()
	fmt.Printf("F5 %s %s\n", r1, r2)
}
