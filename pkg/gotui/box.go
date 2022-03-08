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
	topOrBottom := "-"
	side := "|"

	w := b.w * 2
	var builder strings.Builder
	for i := 0; i < b.h; i++ {
		for j := 0; j < w; j++ {
			if i == 0 && j == 0 {
				builder.WriteString(topLeft)
				builder.WriteString(strings.Repeat(topOrBottom, w-2))
				builder.WriteString(topRight)
				builder.WriteString("\n")
			} else if i == b.h-1 && j == 0 {
				builder.WriteString(bottomLeft)
				builder.WriteString(strings.Repeat(topOrBottom, w-2))
				builder.WriteString(bottomRight)
				builder.WriteString("\n")
			} else if j == 0 {
				builder.WriteString(side)
				builder.WriteString(strings.Repeat(" ", w-2))
				builder.WriteString(side)
				builder.WriteString("\n")
			}

		}
	}
	fmt.Println(builder.String())
}
