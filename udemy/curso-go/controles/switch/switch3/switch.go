package main

import (
	"fmt"
	"time"
)

func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "inteiro"
	case float32, float64:
		return "real"
	case string:
		return "string"
	case func():
		return "função"
	default:
		return "desconhecido"
	}
}

func main() {
	i, f32, f64, s, f, o := 1, 2.3, 2.555555, "Hello", func() {}, time.Now()
	fmt.Println("A variável i é do tipo", tipo(i))
	fmt.Println("A variável f32 é do tipo", tipo(f32))
	fmt.Println("A variável f64 é do tipo", tipo(f64))
	fmt.Println("A variável s é do tipo", tipo(s))
	fmt.Println("A variável f é do tipo", tipo(f))
	fmt.Println("A variável o é do tipo", tipo(o))

}
