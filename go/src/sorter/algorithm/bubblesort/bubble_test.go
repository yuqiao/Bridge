package bubblesort

import (
	"testing"
)

func TestBubbleSort1(t *testing.T) {
	values := []int{5, 4, 3, 2, 1}
	BubbleSort(values)
	if values[0] != 1 || values[1] != 3 {
		t.Error("BubbleSort() failed. Got", values, "Expected  1 2 3 4 5")
	}
}
