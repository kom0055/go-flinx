package flinx

import "testing"

func TestJoin(t *testing.T) {
	outer := []int{0, 1, 2, 3, 4, 5, 8}
	inner := []int{1, 2, 1, 4, 7, 6, 7, 2}
	want := []KeyValue[int, int]{
		{1, 1},
		{1, 1},
		{2, 2},
		{2, 2},
		{4, 4},
	}

	q := Join[int, int, KeyValue[int, int], int](func(i int) int { return i },
		func(i int) int { return i },
		func(outer int, inner int) KeyValue[int, int] {
			return KeyValue[int, int]{outer, inner}
		},
	)(FromSlice[int](outer), FromSlice[int](inner))
	if !validateQuery[KeyValue[int, int]](q, want) {
		t.Errorf("From().Join()=%v expected %v", toSlice(q), want)
	}
}
