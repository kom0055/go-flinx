package flinx

import "testing"

func TestExcept(t *testing.T) {
	input1 := []int{1, 2, 3, 4, 5, 1, 2, 5}
	input2 := []int{1, 2}
	want := []int{3, 4, 5, 5}

	if q := Except[int](FromSlice[int](input1), FromSlice[int](input2)); !validateQuery(q, want) {
		t.Errorf("From(%v).Except(%v)=%v expected %v", input1, input2, toSlice(q), want)
	}
}

func TestExceptBy(t *testing.T) {
	input1 := []int{1, 2, 3, 4, 5, 1, 2, 5}
	input2 := []int{1}
	want := []int{2, 4, 2}

	if q := ExceptBy[int, int](func(i int) int {
		return i % 2
	})(FromSlice[int](input1), FromSlice[int](input2)); !validateQuery(q, want) {
		t.Errorf("From(%v).ExceptBy(%v)=%v expected %v", input1, input2, toSlice(q), want)
	}
}
