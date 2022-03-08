package gotui

import terminal "golang.org/x/term"

type Grid struct {
	Box
	items []*Box
}

const DEFAULT_TERM = 40 // avoid going over 80 characters

func NewGrid() Grid {
	// TODO: get max width of current terminal window
	termW, _, err := terminal.GetSize(1)
	termW = termW / 2
	if err != nil || termW > DEFAULT_TERM {
		termW = DEFAULT_TERM
	}
	box := NewBox(termW, termW, nil)
	grid := Grid{*box, nil}
	return grid
}
