package flinx

import "testing"

func TestZip(t *testing.T) {
	input1 := []int{1, 2, 3}
	input2 := []int{2, 4, 5, 1}
	want := []int{3, 6, 8}

	if q := Zip[int, int, int](func(i, j int) int {
		return i + j
	})(FromSlice[int](input1), FromSlice[int](input2)); !validateQuery(q, want) {
		t.Errorf("From(%v).Zip(%v)=%v expected %v", input1, input2, toSlice(q), want)
	}
}
