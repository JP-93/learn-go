package execicios2

import "testing"

func TestVersion(t *testing.T) {
	got := Version()
	want := "go1.17"

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
