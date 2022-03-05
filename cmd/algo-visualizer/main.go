package main

import (
	"fmt"

	"github.com/jdearly/algo-visualizer/pkg/algo"
)

func main() {
	sorted := algo.InsertionSort([]int{3, 2, 5, 1, 4, 6, 12, 8, 11, 10, 7, 9})
	fmt.Println(sorted)
}
