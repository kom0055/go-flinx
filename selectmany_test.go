package flinx

import (
	"strconv"
	"testing"
)

func TestSelectMany(t *testing.T) {
	{
		tests := []struct {
			input    []string
			selector func(string2 string) Query[rune]
			output   []rune
		}{

			{[]string{"str", "ing"}, func(i string) Query[rune] {
				return FromString(i)
			}, []rune{'s', 't', 'r', 'i', 'n', 'g'}},
		}

		for _, test := range tests {
			if q := SelectMany(test.selector)(FromSlice(test.input)); !ValidateQuery(q, test.output) {
				t.Errorf("From(%v).SelectMany()=%v expected %v", test.input, ToSlice(q), test.output)
			}
		}
	}
	{
		tests := []struct {
			input    [][]int
			selector func([]int) Query[int]
			output   []int
		}{
			{[][]int{{1, 2, 3}, {4, 5, 6, 7}},
				func(i []int) Query[int] {
					return FromSlice(i)
				}, []int{1, 2, 3, 4, 5, 6, 7}},
		}

		for _, test := range tests {
			if q := SelectMany(test.selector)(FromSlice(test.input)); !ValidateQuery(q, test.output) {
				t.Errorf("From(%v).SelectMany()=%v expected %v", test.input, ToSlice(q), test.output)
			}
		}
	}
}

func TestSelectManyIndexed(t *testing.T) {
	{
		tests := []struct {
			input    [][]int
			selector func(int, []int) Query[int]
			output   []int
		}{
			{[][]int{{1, 2, 3}, {4, 5, 6, 7}}, func(i int, x []int) Query[int] {
				if i > 0 {
					return FromSlice(x[1:])
				}
				return FromSlice(x)
			}, []int{1, 2, 3, 5, 6, 7}},
		}

		for _, test := range tests {

			if q := SelectManyIndexed(test.selector)(FromSlice(test.input)); !ValidateQuery(q, test.output) {
				t.Errorf("From(%v).SelectManyIndexed()=%v expected %v", test.input, ToSlice(q), test.output)
			}
		}
	}
	{
		tests := []struct {
			input    []string
			selector func(int, string) Query[rune]
			output   []rune
		}{

			{[]string{"str", "ing"},
				func(i int, x string) Query[rune] {
					return FromString(x + strconv.Itoa(i))
				}, []rune{'s', 't', 'r', '0', 'i', 'n', 'g', '1'}},
		}

		for _, test := range tests {

			if q := SelectManyIndexed(test.selector)(FromSlice(test.input)); !ValidateQuery(q, test.output) {
				t.Errorf("From(%v).SelectManyIndexed()=%v expected %v", test.input, ToSlice(q), test.output)
			}
		}
	}
}

func TestSelectManyBy(t *testing.T) {
	{
		tests := []struct {
			input          [][]int
			selector       func([]int) Query[int]
			resultSelector func(int, []int) int
			output         []int
		}{
			{[][]int{{1, 2, 3}, {4, 5, 6, 7}},
				func(i []int) Query[int] {
					return FromSlice(i)
				}, func(x int, y []int) int {
					return x + 1
				}, []int{2, 3, 4, 5, 6, 7, 8}},
		}

		for _, test := range tests {

			if q := SelectManyBy(test.selector, test.resultSelector)(FromSlice(test.input)); !ValidateQuery(q, test.output) {
				t.Errorf("From(%v).SelectManyBy()=%v expected %v", test.input, ToSlice(q), test.output)
			}
		}
	}
	{
		tests := []struct {
			input          []string
			selector       func(string) Query[rune]
			resultSelector func(rune, string) string
			output         []string
		}{

			{[]string{"str", "ing"},
				func(i string) Query[rune] {
					return FromString(i)
				}, func(x rune, y string) string {
					return string(x) + "_"
				}, []string{"s_", "t_", "r_", "i_", "n_", "g_"}},
		}

		for _, test := range tests {
			if q := SelectManyBy(test.selector, test.resultSelector)(FromSlice(test.input)); !ValidateQuery(q, test.output) {
				t.Errorf("From(%v).SelectManyBy()=%v expected %v", test.input, ToSlice(q), test.output)
			}
		}
	}
}

func TestSelectManyIndexedBy(t *testing.T) {
	{
		tests := []struct {
			input          [][]int
			selector       func(int, []int) Query[int]
			resultSelector func(int, []int) int
			output         []int
		}{
			{[][]int{{1, 2, 3}, {4, 5, 6, 7}},
				func(i int, x []int) Query[int] {
					if i == 0 {
						return FromSlice([]int{10, 20, 30})
					}
					return FromSlice(x)
				}, func(x int, y []int) int {
					return x + 1
				}, []int{11, 21, 31, 5, 6, 7, 8}},
		}

		for _, test := range tests {
			if q := SelectManyByIndexed(test.selector, test.resultSelector)(FromSlice(test.input)); !ValidateQuery(q, test.output) {
				t.Errorf("From(%v).SelectManyIndexedBy()=%v expected %v", test.input, ToSlice(q), test.output)
			}
		}
	}
	{
		tests := []struct {
			input          []string
			selector       func(int, string) Query[rune]
			resultSelector func(rune, string) string
			output         []string
		}{

			{[]string{"st", "ng"},
				func(i int, x string) Query[rune] {
					if i == 0 {
						return FromString(x + "r")
					}
					return FromString("i" + x)
				}, func(x rune, y string) string {
					return string(x) + "_"
				}, []string{"s_", "t_", "r_", "i_", "n_", "g_"}},
		}

		for _, test := range tests {

			if q := SelectManyByIndexed(test.selector, test.resultSelector)(FromSlice(test.input)); !ValidateQuery(q, test.output) {
				t.Errorf("From(%v).SelectManyIndexedBy()=%v expected %v", test.input, ToSlice(q), test.output)
			}
		}
	}
}
