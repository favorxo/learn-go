package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("123")
	want := "Hello, 123"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
