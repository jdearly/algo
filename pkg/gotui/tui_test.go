package gotui

import "testing"

func TestHelloWorld(t *testing.T) {
	expected := "Hello, World!"
	if observed := Tui(); observed != expected {
		t.Fatalf("main() = %v, want %v", observed, expected)
	}
}
