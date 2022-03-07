package algo

import (
	"os"
	"os/exec"
	"time"

	"github.com/jdearly/algo-visualizer/pkg/gotui"
)

func InsertionSort(items []int, output bool) ([]int, int) {
	var j int
	var ops int

	for i := 1; i < len(items); i++ {
		j = i
		for j > 0 && (items[j] < items[j-1]) {
			items[j], items[j-1] = items[j-1], items[j]
			j = j - 1
			ops++
			// print current order
			// TODO: this is temporary until real tui is built

			if output {
				for _, v := range items {
					gotui.PrintStars(v)
				}
				time.Sleep(500 * time.Millisecond)
			}
			cmd := exec.Command("clear")
			cmd.Stdout = os.Stdout
			cmd.Run()
		}
	}

	return items, ops
}
