package flinx

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestDistinct(t *testing.T) {
	{
		tests := []struct {
			input  []int
			output []int
		}{
			{[]int{1, 2, 2, 3, 1}, []int{1, 2, 3}},
			{[]int{1, 1, 1, 2, 1, 2, 3, 4, 2}, []int{1, 2, 3, 4}},
		}

		for _, test := range tests {
			if q := Distinct(FromSlice(test.input)); !ValidateQuery(q, test.output) {
				t.Errorf("From(%v).Distinct()=%v expected %v", test.input, ToSlice(q), test.output)
			}
		}
	}

	{

		input := "sstr"
		output := []rune{'s', 't', 'r'}
		if q := Distinct(FromString(input)); !ValidateQuery(q, output) {
			t.Errorf("From(%v).Distinct()=%v expected %v", input, ToSlice(q), output)
		}
	}

}

func TestDistinctForOrderedQuery(t *testing.T) {
	{
		tests := []struct {
			input  []int
			output []int
		}{
			{[]int{1, 2, 2, 3, 1}, []int{1, 2, 3}},
			{[]int{1, 1, 1, 2, 1, 2, 3, 4, 2}, []int{1, 2, 3, 4}},
		}
		orderByFn := OrderBy(func(i, j int) int {
			return i - j
		}, func(i int) int { return i })
		for _, test := range tests {
			if q := Distinct(FromSlice(test.input)); !ValidateQuery(q, test.output) {
				t.Errorf("From(%v).Distinct()=%v expected %v", test.input, ToSlice(q), test.output)
			}
		}
		for _, test := range tests {

			if q := orderByFn(Distinct(FromSlice(test.input))).Query; !ValidateQuery(q, test.output) {
				t.Errorf("From(%v).Distinct()=%v expected %v", test.input, ToSlice(q), test.output)
			}

		}
	}

	{

		orderByFn := OrderBy(func(i, j rune) int {
			return int(i - j)
		}, func(i rune) rune { return i })

		input := "sstr"
		output := []rune{'r', 's', 't'}

		str := ToString(orderByFn(Distinct(FromString(input))).Query)
		expect := string(output)
		assert.Equal(t, expect, str)

	}

}

func TestDistinctBy(t *testing.T) {
	type user struct {
		id   int
		name string
	}

	users := []user{{1, "Foo"}, {2, "Bar"}, {3, "Foo"}}
	want := []user{{1, "Foo"}, {2, "Bar"}}

	distinctByFn := DistinctBy(func(u user) string {
		return u.name
	})
	if q := distinctByFn(FromSlice(users)); !ValidateQuery(q, want) {
		t.Errorf("From(%v).DistinctBy()=%v expected %v", users, ToSlice(q), want)
	}
}
