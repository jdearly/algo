package gotui

import "testing"

func TestBox(t *testing.T) {
	//	box := NewBox(42, 42, nil)
	//	box.DrawBox()
}
func TestGrid(t *testing.T) {
	grid := NewGrid()
	grid.Box.DrawBox()
}
