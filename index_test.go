package flinx

import (
	"testing"
)

func TestIndexOf(t *testing.T) {

	{
		tests := []struct {
			input     []int
			predicate func(int) bool
			expected  int
		}{
			{
				input: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				predicate: func(i int) bool {
					return i == 3
				},
				expected: 2,
			},
		}

		for _, test := range tests {

			index := IndexOf[int](test.predicate)(FromSlice[int](test.input))
			if index != test.expected {
				t.Errorf("From(%v).IndexOf() expected %v received %v", test.input, test.expected, index)
			}

		}

	}

	{
		tests := []struct {
			input     string
			predicate func(rune) bool
			expected  int
		}{

			{
				input: "sstr",
				predicate: func(i rune) bool {
					return i == 'r'
				},
				expected: 3,
			},
			{
				input: "gadsgsadgsda",
				predicate: func(i rune) bool {
					return i == 'z'
				},
				expected: -1,
			},
		}
		for _, test := range tests {

			index := IndexOf[rune](test.predicate)(FromString(test.input))
			if index != test.expected {
				t.Errorf("From(%v).IndexOf() expected %v received %v", test.input, test.expected, index)
			}

		}

	}

}
