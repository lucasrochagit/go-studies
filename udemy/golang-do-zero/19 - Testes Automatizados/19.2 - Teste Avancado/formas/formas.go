package formas

import (
	"math"
)

type GeometricFigure interface {
	Name() string
	Area() float64
}

type Rectangle struct {
	basis  float64
	height float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) Name() string {
	return "retangulo"
}

func (r Rectangle) Area() float64 {
	return r.basis * r.height
}

func (c Circle) Name() string {
	return "circulo"
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}
