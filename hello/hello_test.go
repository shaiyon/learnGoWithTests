package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("test prefixes", func(t *testing.T) {
		got := helloPrefixes()["English"]
		want := "Hello, "
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello(HelloOptions{name: "Chris"})
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("default language english and name 'World'", func(t *testing.T) {
		got := Hello(HelloOptions{})
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
	
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello(HelloOptions{name: "Elodie", language: "Spanish"})
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("default name in French", func(t *testing.T) {
		got := Hello(HelloOptions{language: "French"})
		want := "Bonjour, World"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
