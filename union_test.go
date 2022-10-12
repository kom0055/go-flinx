package flinx

import "testing"

func TestUnion(t *testing.T) {
	input1 := []int{1, 2, 3}
	input2 := []int{2, 4, 5, 1}
	want := []int{1, 2, 3, 4, 5}

	if q := Union(FromSlice(input1), FromSlice(input2)); !ValidateQuery(q, want) {
		t.Errorf("From(%v).Union(%v)=%v expected %v", input1, input2, ToSlice(q), want)
	}
}
