package flinx

import "testing"

func TestWhere(t *testing.T) {

	{
		tests := []struct {
			input     []int
			predicate func(int) bool
			output    []int
		}{
			{[]int{1, 1, 1, 2, 1, 2, 3, 4, 2}, func(i int) bool {
				return i >= 3
			}, []int{3, 4}},
		}

		for _, test := range tests {

			if q := Where[int](test.predicate)(FromSlice[int](test.input)); !validateQuery(q, test.output) {
				t.Errorf("From(%v).Where()=%v expected %v", test.input, toSlice(q), test.output)
			}
		}
	}

	{
		tests := []struct {
			input     string
			predicate func(rune) bool
			output    []rune
		}{

			{"sstr", func(i rune) bool {
				return i != 's'
			}, []rune{'t', 'r'}},
		}

		for _, test := range tests {

			if q := Where[rune](test.predicate)(FromString(test.input)); !validateQuery(q, test.output) {
				t.Errorf("From(%v).Where()=%v expected %v", test.input, toSlice(q), test.output)
			}
		}
	}

}

func TestWhereIndexed(t *testing.T) {
	{
		tests := []struct {
			input     []int
			predicate func(int, int) bool
			output    []int
		}{
			{[]int{1, 1, 1, 2, 1, 2, 3, 4, 2}, func(i int, x int) bool {
				return x < 4 && i > 4
			}, []int{2, 3, 2}},
		}

		for _, test := range tests {

			if q := WhereIndexed[int](test.predicate)(FromSlice[int](test.input)); !validateQuery(q, test.output) {
				t.Errorf("From(%v).WhereIndexed()=%v expected %v", test.input, toSlice(q), test.output)
			}
		}
	}
	{
		tests := []struct {
			input     string
			predicate func(int, rune) bool
			output    []rune
		}{

			{"sstr", func(i int, x rune) bool {
				return x != 's' || i == 1
			}, []rune{'s', 't', 'r'}},
			{"abcde", func(i int, _ rune) bool {
				return i < 2
			}, []rune{'a', 'b'}},
		}

		for _, test := range tests {

			if q := WhereIndexed[rune](test.predicate)(FromString(test.input)); !validateQuery(q, test.output) {
				t.Errorf("From(%v).WhereIndexed()=%v expected %v", test.input, toSlice(q), test.output)
			}
		}
	}
}
