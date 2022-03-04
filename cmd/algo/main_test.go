package main

import "testing"

func TestHelloWorld(t *testing.T) {
	expected := "Hello, World!"
	if observed := main(); observed != expected {
		t.Fatalf("main() = %v, want %v", observed, expected)
	}
}
