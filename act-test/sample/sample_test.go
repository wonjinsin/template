package sample

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Go")
	want := "Hello, Go!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	got = Hello("Go")
	want = "Hello, Go1!"

	if got == want {
		t.Errorf("got %q want %q", got, want)
	}
}
