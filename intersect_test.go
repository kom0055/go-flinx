package flinx

import "testing"

func TestIntersect(t *testing.T) {
	input1 := []int{1, 2, 3}
	input2 := []int{1, 4, 7, 9, 12, 3}
	want := []int{1, 3}

	if q := Intersect(FromSlice(input1), FromSlice(input2)); !ValidateQuery(q, want) {
		t.Errorf("From(%v).Intersect(%v)=%v expected %v", input1, input2, ToSlice(q), want)
	}
}

func TestIntersectBy(t *testing.T) {
	input1 := []int{5, 7, 8}
	input2 := []int{1, 4, 7, 9, 12, 3}
	want := []int{5, 8}

	if q := IntersectBy(func(i int) int {
		return i % 2
	})(FromSlice(input1), FromSlice(input2)); !ValidateQuery(q, want) {
		t.Errorf("From(%v).IntersectBy(%v)=%v expected %v", input1, input2, ToSlice(q), want)
	}
}
