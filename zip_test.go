package flinx

import "testing"

func TestZip(t *testing.T) {
	input1 := []int{1, 2, 3}
	input2 := []int{2, 4, 5, 1}
	want := []int{3, 6, 8}

	if q := Zip(func(i, j int) int {
		return i + j
	})(FromSlice(input1), FromSlice(input2)); !ValidateQuery(q, want) {
		t.Errorf("From(%v).Zip(%v)=%v expected %v", input1, input2, ToSlice(q), want)
	}
}
