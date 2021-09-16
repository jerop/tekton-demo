package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Earth")
	want := "Hello Earth!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}