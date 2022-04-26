package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Sam")
	want := "Hello, Sam"

	if want != got {
		t.Errorf("got %q, want %q", got, want)
	}
}
