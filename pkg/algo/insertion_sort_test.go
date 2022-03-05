package algo

import (
	"reflect"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	expected := []int{1, 2, 3, 4, 5}
	if observed := InsertionSort([]int{4, 1, 3, 5, 2}); !reflect.DeepEqual(observed, expected) {
		t.Fatalf("main() = %v, want %v", observed, expected)
	}
}
