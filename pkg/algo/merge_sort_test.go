package algo

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	if observed := MergeSort([]int{4, 2, 7, 5, 8, 1, 3, 9, 6}); !reflect.DeepEqual(observed, expected) {
		t.Fatalf("MergeSort() = %v, want %v", observed, expected)
	}
}
