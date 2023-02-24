package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("George!")
	want := "Hello, George!"

	if got != want {
		t.Errorf("got %s want %q", got, want)
	}
}