package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %s want %q", got, want)
		}
	}
	t.Run("saying hello", func(t *testing.T) {
		got := Hello("George!", "")
		want := "Hello, George!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'hello world' when an empty string is supplied ", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("George", spanish)
		want := "Hola, George"

		assertCorrectMessage(t, got, want)
	})
}
