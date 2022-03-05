package gotui

import "testing"

func TestHelloWorld(t *testing.T) {
	expected := "*****"
	if observed := PrintStars(5); observed != expected {
		t.Fatalf("PrintStars() = %v, want %v", observed, expected)
	}
}
