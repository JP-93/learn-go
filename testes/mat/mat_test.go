package mat

import "testing"

const erroPadrao = "Valor esperado %v, valor retornado %v"

func TestMedia(t *testing.T) {
	v1 := 7.28
	valor := Media(7.2)

	if valor != v1 {
		t.Errorf(erroPadrao, v1, valor)
	}
}
