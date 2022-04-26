package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, want, got string) {
		t.Helper()
		if want != got {
			t.Errorf("want %q, got %q", want, got)
		}
	}
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Sam")
		want := "Hello, Sam"
		assertCorrectMessage(t, want, got)
	})
	t.Run("say 'Hello, World' when an emtpy string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, world"
		assertCorrectMessage(t, want, got)
	})
}
