package flinx

import "testing"

func TestAppend(t *testing.T) {
	input := []int{1, 2, 3, 4}
	want := []int{1, 2, 3, 4, 5}
	if q := Append(FromSlice(input), 5); !ValidateQuery(q, want) {
		t.Errorf("From(%v).Append()=%v expected %v", input, ToSlice(q), want)
	}
}

func TestConcat(t *testing.T) {
	input1 := []int{1, 2, 3}
	input2 := []int{4, 5}
	want := []int{1, 2, 3, 4, 5}

	if q := Concat(FromSlice(input1), FromSlice(input2)); !ValidateQuery(q, want) {
		t.Errorf("From(%v).Concat(%v)=%v expected %v", input1, input2, ToSlice(q), want)
	}
}

func TestPrepend(t *testing.T) {
	input := []int{1, 2, 3, 4}
	want := []int{0, 1, 2, 3, 4}
	if q := Prepend(FromSlice(input), 0); !ValidateQuery(q, want) {
		t.Errorf("From(%v).Prepend()=%v expected %v", input, ToSlice(q), want)
	}
}
