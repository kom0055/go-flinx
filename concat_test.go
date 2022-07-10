package flinx

import "testing"

func TestAppend(t *testing.T) {
	input := []int{1, 2, 3, 4}
	want := []int{1, 2, 3, 4, 5}
	appendFn := Append[int](5)
	if q := appendFn(FromSlice[int](input)); !validateQuery[int](q, want) {
		t.Errorf("From(%v).Append()=%v expected %v", input, toSlice[int](q), want)
	}
}

func TestConcat(t *testing.T) {
	input1 := []int{1, 2, 3}
	input2 := []int{4, 5}
	want := []int{1, 2, 3, 4, 5}

	concatFn := Concat[int]
	if q := concatFn(FromSlice[int](input1), FromSlice[int](input2)); !validateQuery[int](q, want) {
		t.Errorf("From(%v).Concat(%v)=%v expected %v", input1, input2, toSlice(q), want)
	}
}

func TestPrepend(t *testing.T) {
	input := []int{1, 2, 3, 4}
	want := []int{0, 1, 2, 3, 4}
	prependFn := Prepend[int](0)
	if q := prependFn(FromSlice[int](input)); !validateQuery[int](q, want) {
		t.Errorf("From(%v).Prepend()=%v expected %v", input, toSlice(q), want)
	}
}
