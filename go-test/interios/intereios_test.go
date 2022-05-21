package inteiros

import "testing"

func TestS(t *testing.T) {
	got := Soma(2, 2)
	want := 4

	if got != want {
		t.Errorf("expected '%d' but got '%d' ", got, want)
	}

}
