package algo

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	operations := 36
	if observed, ops := BubbleSort([]int{4, 2, 7, 5, 8, 1, 3, 9, 6}, false); !reflect.DeepEqual(observed, expected) || ops != operations {
		t.Fatalf("BubbleSort() = %v, want %v", observed, expected)
	}
}
