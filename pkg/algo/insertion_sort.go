package main

func InsertionSort(items []int, n int) []int {
	var j int

	for i := 1; i < n; i++ {
		j = i
		for j > 0 && (items[j] < items[j-1]) {
			items[j], items[j-1] = items[j-1], items[j]
			j = j - 1
		}
	}

	return items
}
