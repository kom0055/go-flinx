package flinx

import (
	"gotest.tools/v3/assert"
	"testing"
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
			if q := Distinct[int](FromSlice[int](test.input)); !validateQuery[int](q, test.output) {
				t.Errorf("From(%v).Distinct()=%v expected %v", test.input, toSlice(q), test.output)
			}
		}
	}

	{

		input := "sstr"
		output := []rune{'s', 't', 'r'}
		if q := Distinct[rune](FromString(input)); !validateQuery[rune](q, output) {
			t.Errorf("From(%v).Distinct()=%v expected %v", input, toSlice(q), output)
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
		distinctFn := Distinct[int]
		orderByFn := OrderBy[int](func(i, j int) int {
			return i - j
		}, func(i int) int { return i })
		for _, test := range tests {
			if q := distinctFn(FromSlice[int](test.input)); !validateQuery[int](q, test.output) {
				t.Errorf("From(%v).Distinct()=%v expected %v", test.input, toSlice(q), test.output)
			}
		}
		for _, test := range tests {

			if q := orderByFn(distinctFn(FromSlice[int](test.input))).Query; !validateQuery[int](q, test.output) {
				t.Errorf("From(%v).Distinct()=%v expected %v", test.input, toSlice(q), test.output)
			}

		}
	}

	{

		distinctFn := Distinct[rune]
		orderByFn := OrderBy[rune](func(i, j rune) int {
			return int(i - j)
		}, func(i rune) rune { return i })

		input := "sstr"
		output := []rune{'r', 's', 't'}

		str := ToString(orderByFn(distinctFn(FromString(input))).Query)
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
	want := []user{user{1, "Foo"}, user{2, "Bar"}}

	distinctByFn := DistinctBy[user, string](func(u user) string {
		return u.name
	})
	if q := distinctByFn(FromSlice[user](users)); !validateQuery[user](q, want) {
		t.Errorf("From(%v).DistinctBy()=%v expected %v", users, toSlice(q), want)
	}
}
