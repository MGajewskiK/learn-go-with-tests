package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say in Spanish", func(t *testing.T) {
		got := Hello("Elondie", "Spanish")
		want := "Hola, Elondie"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say in French", func(t *testing.T) {
		got := Hello("Elondie", "French")
		want := "Bonjour, Elondie"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say in Polish", func(t *testing.T) {
		got := Hello("Elondie", "Polish")
		want := "Witaj, Elondie"

		assertCorrectMessage(t, got, want)
	})
}
