package formas

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	t.Run("Retângulo", func(t *testing.T) {
		ret := Rectangle{10, 12}
		expected := float64(120)
		result := ret.Area()

		if expected != result {
			t.Fatalf("Area esperada (%f) eh diferente da esperada (%f)", result, expected)
		}
	})

	t.Run("Círculo", func(t *testing.T) {
		circ := Circle{10}
		expected := float64(math.Pi * math.Pow(10, 2))
		result := circ.Area()

		if expected != result {
			t.Fatalf("Area esperada (%f) eh diferente da esperada (%f)", result, expected)
		}
	})
}
