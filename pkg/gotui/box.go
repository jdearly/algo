package gotui

import (
	"fmt"
	"strings"
)

type Box struct {
	w      int
	h      int
	parent *Box
}

func NewBox(w, h int, parent *Box) *Box {
	return &Box{w, h, parent}
}

func (b *Box) DrawBox() {
	// TODO: do this better?
	topLeft := "┏"
	topRight := "┓"
	bottomLeft := "┗"
	bottomRight := "┛"
	topOrBottom := "━"
	side := "┃"

	w := b.w * 2
	for i := 0; i < b.h; i++ {
		for j := 0; j < w; j++ {
			if i == 0 && j == 0 {
				fmt.Println(topLeft + strings.Repeat(topOrBottom, w-2) + topRight)
			} else if i == b.h-1 && j == 0 {
				fmt.Println(bottomLeft + strings.Repeat(topOrBottom, w-2) + bottomRight)
			} else if j == 0 {
				fmt.Println(side + strings.Repeat(" ", w-2) + side)
			}

		}
	}

}
