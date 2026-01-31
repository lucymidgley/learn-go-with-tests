package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Lucy")
		want := "Hello, Lucy"
	
		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello to without a name supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello,j world"
	
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
