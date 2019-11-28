package hello

import "testing"

func TestHello(t *testing.T) {
    want := "Hello , world."
    if got := SayHello(", world."); got != want {
        t.Errorf("SayHello() = %q, want %q", got, want)
    }
}