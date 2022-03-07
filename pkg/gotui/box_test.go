package gotui

import "testing"

func TestBox(t *testing.T) {
	box := NewBox(45, 45, nil)
	box.DrawBox()
}
