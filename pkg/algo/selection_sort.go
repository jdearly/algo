package algo

func SelectionSort(items []int, output bool) ([]int, int) {
	var n = len(items)
	var swaps int
	for i := 0; i < n; i++ {
		var min = i
		for j := i; j < n; j++ {
			if items[j] < items[min] {
				min = j
			}
		}
		items[i], items[min] = items[min], items[i]
		swaps++
	}
	return items, swaps
}
