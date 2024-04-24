package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("should return Hola to people when name is not empty and language is spanish", func(t *testing.T) {
		name := "Satya"
		got := Hello(name, "Spanish")
		want := "Hola Satya"
		if got != want {
			t.Errorf("want %q got %q", want, got)
		}
	})

	t.Run("should return Bonjour to people when name is not empty and language is spanish", func(t *testing.T) {
		name := "Satya"
		got := Hello(name, "French")
		want := "Bonjour Satya"
		if got != want {
			t.Errorf("want %q got %q", want, got)
		}
	})

	t.Run("should return Hello World when name and language is passed an empty string", func(t *testing.T) {
		name := ""
		got := Hello(name, "")
		want := "Hello World"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t *testing.T, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("want %q got %q", want, got)
	}
}
