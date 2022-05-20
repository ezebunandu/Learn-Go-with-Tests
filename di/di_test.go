package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}

	Greet(&buffer, "Chris")

	want := "Hello, Chris"
	got := buffer.String()

	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}
