package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Banana == Banana", "Banana" == "Banana")
	fmt.Println("3 != 2", 3 != 2)
	fmt.Println("3 < 2", 3 < 2)
	fmt.Println("3 > 2", 3 > 2)
	fmt.Println("3 <= 2", 3 <= 2)
	fmt.Println("3 >= 2", 3 >= 2)

	d1, d2 := time.Unix(0, 0), time.Unix(0, 0)
	fmt.Println("d1 == d2", d1 == d2)        // true
	fmt.Println("d1 equal d2", d1.Equal(d2)) // true

	type Pessoa struct {
		Nome string
	}

	p1, p2 := Pessoa{"João"}, Pessoa{"João"}
	fmt.Println("p1 == p2", p1 == p2)
}
