package uno

import "testing"


func TestHello(t *testing.T) {
    want := "Hi, world !!"
    if got := Hi("world"); got != want {
        t.Errorf("Hello() = %q, want %q", got, want)
    }
}
