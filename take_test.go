package flinx

import "testing"

func TestTake(t *testing.T) {
	{
		tests := []struct {
			input  []int
			output []int
		}{
			{[]int{1, 2, 2, 3, 1}, []int{1, 2, 2}},
			{[]int{1, 1, 1, 2, 1, 2, 3, 4, 2}, []int{1, 1, 1}},
		}

		for _, test := range tests {

			if q := Take(FromSlice(test.input), 3); !ValidateQuery(q, test.output) {
				t.Errorf("From(%v).Take(3)=%v expected %v", test.input, ToSlice(q), test.output)
			}
		}
	}
	{
		tests := []struct {
			input  string
			output []rune
		}{

			{"sstr", []rune{'s', 's', 't'}},
		}

		for _, test := range tests {
			if q := Take(FromString(test.input), 3); !ValidateQuery(q, test.output) {
				t.Errorf("From(%v).Take(3)=%v expected %v", test.input, ToSlice(q), test.output)
			}
		}
	}
}

func TestTakeWhile(t *testing.T) {

	{
		tests := []struct {
			input     []int
			predicate func(int2 int) bool
			output    []int
		}{
			{[]int{1, 1, 1, 2, 1, 2}, func(i int) bool {
				return i < 3
			}, []int{1, 1, 1, 2, 1, 2}},
			{[]int{1, 1, 1, 2, 1, 2, 3, 4, 2}, func(i int) bool {
				return i < 3
			}, []int{1, 1, 1, 2, 1, 2}},
		}

		for _, test := range tests {

			if q := TakeWhile(test.predicate)(FromSlice(test.input)); !ValidateQuery(q, test.output) {
				t.Errorf("From(%v).Take(3)=%v expected %v", test.input, ToSlice(q), test.output)
			}
		}
	}
	{
		tests := []struct {
			input     string
			predicate func(int2 rune) bool
			output    []rune
		}{

			{"sstr", func(i rune) bool {
				return i == 's'
			}, []rune{'s', 's'}},
		}
		for _, test := range tests {

			if q := TakeWhile(test.predicate)(FromString(test.input)); !ValidateQuery(q, test.output) {
				t.Errorf("From(%v).Take(3)=%v expected %v", test.input, ToSlice(q), test.output)
			}
		}
	}
}

func TestTakeWhileIndexed(t *testing.T) {

	{
		tests := []struct {
			input     []int
			predicate func(int, int) bool
			output    []int
		}{
			{[]int{1, 1, 1, 2}, func(i int, x int) bool {
				return x < 2 || i < 5
			}, []int{1, 1, 1, 2}},
			{[]int{1, 1, 1, 2, 1, 2, 3, 4, 2}, func(i int, x int) bool {
				return x < 2 || i < 5
			}, []int{1, 1, 1, 2, 1}},
		}

		for _, test := range tests {

			if q := TakeWhileIndexed(test.predicate)(FromSlice(test.input)); !ValidateQuery(q, test.output) {
				t.Errorf("From(%v).Take(3)=%v expected %v", test.input, ToSlice(q), test.output)
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
				return x == 's' && i < 1
			}, []rune{'s'}},
		}

		for _, test := range tests {

			if q := TakeWhileIndexed(test.predicate)(FromString(test.input)); !ValidateQuery(q, test.output) {
				t.Errorf("From(%v).Take(3)=%v expected %v", test.input, ToSlice(q), test.output)
			}
		}
	}
}
