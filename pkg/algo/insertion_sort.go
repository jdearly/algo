package algo

import (
	"fmt"
	"time"

	"github.com/jdearly/algo-visualizer/pkg/gotui"
)

func InsertionSort(items []int) []int {
	var j int

	for i := 1; i < len(items); i++ {
		j = i
		for j > 0 && (items[j] < items[j-1]) {
			items[j], items[j-1] = items[j-1], items[j]
			j = j - 1

			// print current order
			// TODO: this is temporary until real tui is built
			for i, v := range items {
				if i == j {
					fmt.Println(gotui.PrintStars(v) + " <- j-1")
				} else if i == j+1 {
					fmt.Println(gotui.PrintStars(v) + " <- j")
				} else {
					fmt.Println(gotui.PrintStars(v))
				}
			}
			time.Sleep(1 * time.Second)
		}
	}

	return items
}
