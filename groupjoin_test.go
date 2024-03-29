package flinx

import "testing"

func TestGroupJoin(t *testing.T) {
	outer := []int{0, 1, 2}
	inner := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	want := []KeyValue[int, int]{
		{0, 4},
		{1, 5},
		{2, 0},
	}

	q := GroupJoin(
		func(t int) int {
			return t
		},
		func(t int) int {
			return t % 2
		},
		func(outer int, inners []int) KeyValue[int, int] {
			return KeyValue[int, int]{
				outer, len(inners),
			}
		},
	)(FromSlice(outer), FromSlice(inner))

	if !ValidateQuery(q, want) {
		t.Errorf("From().GroupJoin()=%v expected %v", ToSlice(q), want)
	}
}
