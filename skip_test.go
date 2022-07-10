package flinx

import (
	"gotest.tools/v3/assert"
	"testing"
)

func TestSkip(t *testing.T) {

	{
		tests := []struct {
			input  []int
			output []int
		}{
			{[]int{1, 2}, []int{}},
			{[]int{1, 2, 2, 3, 1}, []int{3, 1}},
			{[]int{1, 1, 1, 2, 1, 2, 3, 4, 2}, []int{2, 1, 2, 3, 4, 2}},
		}

		for _, test := range tests {
			res := ToSlice[int](Skip[int](3)(FromSlice[int](test.input)))
			assert.DeepEqual(t, res, test.output)

		}
	}
	{
		tests := []struct {
			input  string
			output []rune
		}{

			{"sstr", []rune{'r'}},
		}

		for _, test := range tests {
			res := ToSlice[rune](Skip[rune](3)(FromString(test.input)))
			assert.DeepEqual(t, res, test.output)

		}
	}
}

func TestSkipWhile(t *testing.T) {

	{
		tests := []struct {
			input     []int
			predicate func(int) bool
			output    []int
		}{
			{[]int{1, 2}, func(i int) bool {
				return i < 3
			}, []int{}},
			{[]int{4, 1, 2}, func(i int) bool {
				return i < 3
			}, []int{4, 1, 2}},
			{[]int{1, 1, 1, 2, 1, 2, 3, 4, 2}, func(i int) bool {
				return i < 3
			}, []int{3, 4, 2}},
		}

		for _, test := range tests {

			if q := SkipWhile[int](test.predicate)(FromSlice[int](test.input)); !validateQuery(q, test.output) {
				t.Errorf("From(%v).SkipWhile()=%v expected %v", test.input, toSlice(q), test.output)
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
				return i == 's'
			}, []rune{'t', 'r'}},
		}

		for _, test := range tests {
			if q := SkipWhile[rune](test.predicate)(FromString(test.input)); !validateQuery(q, test.output) {
				t.Errorf("From(%v).SkipWhile()=%v expected %v", test.input, toSlice(q), test.output)
			}
		}
	}

}

func TestSkipWhileIndexed(t *testing.T) {
	{
		tests := []struct {
			input     []int
			predicate func(int, int) bool
			output    []int
		}{
			{[]int{1, 2}, func(i int, x int) bool {
				return x < 3
			}, []int{}},
			{[]int{4, 1, 2}, func(i int, x int) bool {
				return x < 3
			}, []int{4, 1, 2}},
			{[]int{1, 1, 1, 2, 1, 2, 3, 4, 2}, func(i int, x int) bool {
				return x < 2 || i < 5
			}, []int{2, 3, 4, 2}},
		}

		for _, test := range tests {

			if q := SkipWhileIndexed[int](test.predicate)(FromSlice[int](test.input)); !validateQuery(q, test.output) {
				t.Errorf("From(%v).SkipWhileIndexed()=%v expected %v", test.input, toSlice(q), test.output)
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
			}, []rune{'s', 't', 'r'}},
		}

		for _, test := range tests {

			if q := SkipWhileIndexed[rune](test.predicate)(FromString(test.input)); !validateQuery(q, test.output) {
				t.Errorf("From(%v).SkipWhileIndexed()=%v expected %v", test.input, toSlice(q), test.output)
			}
		}
	}
}
