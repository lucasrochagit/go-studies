package main

import (
	"fmt"
	"math"
)

type GeometricFigure interface {
	name() string
	area() float64
}

func calculateArea(g GeometricFigure) {
	fmt.Printf("A area da forma geométrica %s é %0.2f\n", g.name(), g.area())
}

type Rectangle struct {
	basis  float64
	height float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) name() string {
	return "retangulo"
}

func (r Rectangle) area() float64 {
	return r.basis * r.height
}

func (c Circle) name() string {
	return "circulo"
}

func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func main() {
	r := Rectangle{10, 15}
	calculateArea(r)

	c := Circle{10}
	calculateArea(c)
}
