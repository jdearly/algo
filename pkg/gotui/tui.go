package gotui

import (
	"fmt"
	"strings"
)

func PrintStars(val int) string {
	return fmt.Sprint(strings.Repeat("*", val))
}
