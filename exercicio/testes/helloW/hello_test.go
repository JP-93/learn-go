package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Ola Mundo")
	want := "Ola Mundo"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
