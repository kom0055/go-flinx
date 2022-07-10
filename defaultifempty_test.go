package flinx

import (
	"testing"
)

func TestDefaultIfEmpty(t *testing.T) {
	defaultValue := 0
	tests := []struct {
		input []int
		want  []int
	}{
		{[]int{}, []int{defaultValue}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
	}

	defaltIfEmptyFn := DefaultIfEmpty[int](0)
	for _, test := range tests {
		q := defaltIfEmptyFn(FromSlice(test.input))

		if !validateQuery(q, test.want) {
			t.Errorf("From(%v).DefaultIfEmpty(%v)=%v expected %v", test.input, defaultValue, toSlice[int](q), test.want)
		}
	}

}
