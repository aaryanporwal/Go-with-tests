package main

import "testing"

//  Using "Subtests" to group tests around a "thing" and then have subtests describing diff scenarios

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Aaryan")
		want := "Hello, Aaryan!"

		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World!' when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World!"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}