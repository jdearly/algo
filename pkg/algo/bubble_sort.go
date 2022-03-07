package algo

import (
	"os"
	"os/exec"
	"time"

	"github.com/jdearly/algo-visualizer/pkg/gotui"
)

func BubbleSort(items []int, output bool) ([]int, int) {
	var ops int
	for i := 0; i < len(items)-1; i++ {
		for j := 0; j < len(items)-i-1; j++ {
			if items[j] > items[j+1] {
				items[j], items[j+1] = items[j+1], items[j]
			}
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
