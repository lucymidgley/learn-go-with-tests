package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Lucy", "")
		want := "Hello, Lucy"

		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello to without a name supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"

		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello in Spanish", func(t *testing.T) {
		got := Hello("Alice", "sp")
		want := "Hola, Alice"

		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello in Japanese", func(t *testing.T) {
		got := Hello("Alice", "jp")
		want := "こんにちは, Alice"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
