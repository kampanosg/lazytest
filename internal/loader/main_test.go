package loader

import "testing"

func TestHello(t *testing.T) {
	got := "Hello, Chris"
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestGoodbye(t *testing.T) {
	got := "Goodbye, Chris"
	want := "Goodbye, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
