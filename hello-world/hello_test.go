package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Lucy")
	want := "Hello, Lucy"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
