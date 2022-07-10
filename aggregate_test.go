package flinx

import (
	"gotest.tools/v3/assert"
	"testing"
)
import "strings"

func Test_Aggregate(t *testing.T) {
	First[int](Select[int, int](func(i int) int {
		return i * 2
	})(Where[int](func(i int) bool {
		return i%2 == 0
	})(Range(1, 10))))

	tests := []struct {
		input []string
		want  interface{}
	}{
		{[]string{"apple", "mango", "orange", "passionfruit", "grape"}, "passionfruit"},
		{[]string{}, ""},
	}

	aggr := Aggregate[string](func(r, i string) string {
		if len(r) > len(i) {
			return r
		}
		return i
	})
	for _, test := range tests {

		r := aggr(FromSlice[string](test.input))
		assert.Equal(t, r, test.want)

	}
}

func TestAggregateWithSeed(t *testing.T) {
	input := []string{"apple", "mango", "orange", "banana", "grape"}
	want := "passionfruit"

	aggr := AggregateWithSeed[string](want, func(r, i string) string {
		if len(r) > len(i) {
			return r
		}
		return i
	})
	r := aggr(FromSlice[string](input))

	assert.Equal(t, r, want)
}

func TestAggregateWithSeedBy(t *testing.T) {
	input := []string{"apple", "mango", "orange", "passionfruit", "grape"}
	want := "PASSIONFRUIT"

	aggr := AggregateWithSeedBy[string, string]("banana", func(r, i string) string {
		if len(r) > len(i) {
			return r
		}
		return i
	},
		func(r string) string {
			return strings.ToUpper(r)
		})
	r := aggr(FromSlice[string](input))

	assert.Equal(t, r, want)
}
