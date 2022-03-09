package algo

func MergeSort(items []int) []int {
	if len(items) < 2 {
		return items
	}

	mid := len(items) / 2
	a := MergeSort(items[:mid])
	b := MergeSort(items[mid:])

	return merge(a, b)
}

func merge(a, b []int) []int {
	m := make([]int, len(a)+len(b))
	i, j := 0, 0

	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			m[i+j] = a[i]
			i++
		} else {
			m[i+j] = b[j]
			j++
		}
	}

	for i < len(a) {
		m[i+j] = a[i]
		i++
	}
	for j < len(b) {
		m[i+j] = b[j]
		j++
	}

	return m
}
