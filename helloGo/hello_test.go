package main

import "testing"

func TestHelloWorld(t *testing.T) {
	got := Hello("Tanit")
	want := "Hello, Tanit"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
