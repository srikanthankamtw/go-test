package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello World!"
	if got != want {
		t.Errorf("want %q got %q", want, got)
	}
}
