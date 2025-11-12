package hello

import (
	"testing"
)

func TestHello(t *testing.T) {
	got := Hello("Me")
	want := "hello Me"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
