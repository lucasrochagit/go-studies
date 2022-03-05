package strings

import (
	"strings"
	"testing"
)

const msgIndex = "%s (parte: %s) - índices: esperando (%d) <> encontrado (%d)"

func TestIndex(t *testing.T) {
	t.Parallel()
	tests := []struct {
		texto    string
		parte    string
		esperado int
	}{
		{"Cod3r é show", "Cod3r", 0},
		{"", "", 0},
		{"Opa", "opa", -1},
		{"leonardo", "o", 2},
	}

	for _, test := range tests {
		t.Logf("Massa: %v", test)
		atual := strings.Index(test.texto, test.parte)

		if atual != test.esperado {
			t.Errorf(msgIndex, test.texto, test.parte, test.esperado, atual)
		}
	}
}
