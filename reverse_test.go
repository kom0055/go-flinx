package flinx

import "testing"

func TestReverse(t *testing.T) {
	tests := []struct {
		input []int
		want  []int
	}{
		{[]int{1, 2, 3}, []int{3, 2, 1}},
	}

	for _, test := range tests {

		if q := Reverse(FromSlice(test.input)); !ValidateQuery(q, test.want) {
			t.Errorf("From(%v).Reverse()=%v expected %v", test.input, ToSlice(q), test.want)
		}
	}
}
