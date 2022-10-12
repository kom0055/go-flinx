package flinx

import (
	"strconv"
	"testing"
)

func TestSelect(t *testing.T) {
	{
		tests := []struct {
			input    []int
			selector func(int) int
			output   []int
		}{
			{[]int{1, 2, 3}, func(i int) int {
				return i * 2
			}, []int{2, 4, 6}},
		}

		for _, test := range tests {

			if q := Select(test.selector)(FromSlice(test.input)); !ValidateQuery(q, test.output) {
				t.Errorf("From(%v).Select()=%v expected %v", test.input, ToSlice(q), test.output)
			}
		}
	}
	{
		tests := []struct {
			input    string
			selector func(rune) string
			output   []string
		}{

			{"str", func(i rune) string {
				return string(i) + "1"
			}, []string{"s1", "t1", "r1"}},
		}

		for _, test := range tests {

			if q := Select(test.selector)(FromString(test.input)); !ValidateQuery(q, test.output) {
				t.Errorf("From(%v).Select()=%v expected %v", test.input, ToSlice(q), test.output)
			}
		}
	}
}

func TestSelectIndexed(t *testing.T) {
	{
		tests := []struct {
			input    []int
			selector func(int, int) int
			output   []int
		}{
			{[]int{1, 2, 3}, func(i int, x int) int {
				return x * i
			}, []int{0, 2, 6}},
		}

		for _, test := range tests {

			if q := SelectIndexed(test.selector)(FromSlice(test.input)); !ValidateQuery(q, test.output) {
				t.Errorf("From(%v).SelectIndexed()=%v expected %v", test.input, ToSlice(q), test.output)
			}
		}
	}
	{
		tests := []struct {
			input    string
			selector func(int, rune) string
			output   []string
		}{

			{"str", func(i int, x rune) string {
				return string(x) + strconv.Itoa(i)
			}, []string{"s0", "t1", "r2"}},
		}

		for _, test := range tests {

			if q := SelectIndexed(test.selector)(FromString(test.input)); !ValidateQuery(q, test.output) {
				t.Errorf("From(%v).SelectIndexed()=%v expected %v", test.input, ToSlice(q), test.output)
			}
		}
	}
}
