package flinx

import (
	"testing"

	"gotest.tools/v3/assert"
)
import "strings"

func Test_Aggregate(t *testing.T) {
	First(Select(func(i int) int {
		return i * 2
	})(Where(func(i int) bool {
		return i%2 == 0
	})(Range(1, 10))))

	tests := []struct {
		input []string
		want  interface{}
	}{
		{[]string{"apple", "mango", "orange", "passionfruit", "grape"}, "passionfruit"},
		{[]string{}, ""},
	}

	aggr := Aggregate(func(r, i string) string {
		if len(r) > len(i) {
			return r
		}
		return i
	})
	for _, test := range tests {

		r := aggr(FromSlice(test.input))
		assert.Equal(t, r, test.want)

	}
}

func TestAggregateWithSeed(t *testing.T) {
	input := []string{"apple", "mango", "orange", "banana", "grape"}
	want := "passionfruit"

	aggr := AggregateWithSeed(want, func(r, i string) string {
		if len(r) > len(i) {
			return r
		}
		return i
	})
	r := aggr(FromSlice(input))

	assert.Equal(t, r, want)
}

func TestAggregateWithSeedBy(t *testing.T) {
	input := []string{"apple", "mango", "orange", "passionfruit", "grape"}
	want := "PASSIONFRUIT"

	aggr := AggregateWithSeedBy("banana", func(r, i string) string {
		if len(r) > len(i) {
			return r
		}
		return i
	},
		func(r string) string {
			return strings.ToUpper(r)
		})
	r := aggr(FromSlice(input))

	assert.Equal(t, r, want)
}
