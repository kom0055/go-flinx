package flinx

import (
	"reflect"
	"testing"
)

func TestGroupBy(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	wantEven := []int{2, 4, 6, 8}
	wantOdd := []int{1, 3, 5, 7, 9}

	q := GroupBy(
		func(t int) int {
			return t % 2
		},
		func(t int) int {
			return t
		},
	)(FromSlice(input))

	next := q.Iterate()
	eq := true
	for item, ok := next(); ok; item, ok = next() {
		group := item
		switch group.Key {
		case 0:
			if !reflect.DeepEqual(group.Group, wantEven) {
				eq = false
			}
		case 1:
			if !reflect.DeepEqual(group.Group, wantOdd) {
				eq = false
			}
		default:
			eq = false
		}
	}

	if !eq {
		t.Errorf("From(%v).GroupBy()=%v", input, ToSlice(q))
	}
}
