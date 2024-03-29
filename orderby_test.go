package flinx

import (
	"testing"

	"github.com/kom0055/go-flinx/generics"
)

func TestEmpty(t *testing.T) {
	q := OrderBy(generics.NumericCompare[int], func(in string) int {
		return 0
	})(FromSlice([]string{}))

	_, ok := q.Iterate()()
	if ok {
		t.Errorf("Iterator for empty collection must return ok=false")
	}
}

func TestOrderBy(t *testing.T) {
	slice := make([]foo, 100)

	for i := len(slice) - 1; i >= 0; i-- {
		slice[i].f1 = i
	}

	q := OrderBy(generics.NumericCompare[int], getF1)(FromSlice(slice))

	j := 0
	next := q.Iterate()
	for item, ok := next(); ok; item, ok = next() {
		if item.f1 != j {
			t.Errorf("OrderBy()[%v]=%v expected %v", j, item, foo{f1: j})
		}

		j++
	}
}

func TestOrderByDescending(t *testing.T) {
	slice := make([]foo, 100)

	for i := 0; i < len(slice); i++ {
		slice[i].f1 = i
	}

	q := OrderByDescending(generics.NumericCompare[int],
		getF1)(FromSlice(slice))

	j := len(slice) - 1
	next := q.Iterate()
	for item, ok := next(); ok; item, ok = next() {
		if item.f1 != j {
			t.Errorf("OrderByDescending()[%v]=%v expected %v", j, item, foo{f1: j})
		}

		j--
	}
}

func TestThenBy(t *testing.T) {
	slice := make([]foo, 1000)

	for i := len(slice) - 1; i >= 0; i-- {
		slice[i].f1 = i
		slice[i].f2 = i%2 == 0
	}

	q := ThenBy(generics.OrderedCompare[int], getF1)(
		OrderBy(generics.BoolCompare, getF2)(FromSlice(slice)),
	)

	next := q.Iterate()
	for item, ok := next(); ok; item, ok = next() {
		if item.f2 != (item.f1%2 == 0) {
			t.Errorf("OrderBy().ThenBy()=%v", item)
		}
	}
}

func TestThenByDescending(t *testing.T) {
	slice := make([]foo, 1000)

	for i := len(slice) - 1; i >= 0; i-- {
		slice[i].f1 = i
		slice[i].f2 = i%2 == 0
	}

	orderByFn := OrderBy(generics.BoolCompare, getF2)(FromSlice(slice))
	thenByDescFn := ThenByDescending(generics.NumericCompare[int], getF1)(orderByFn)
	q := thenByDescFn

	next := q.Iterate()
	for item, ok := next(); ok; item, ok = next() {
		if item.f2 != (item.f1%2 == 0) {
			t.Errorf("OrderBy().ThenByDescending()=%v", item)
		}
	}
}

func TestSort(t *testing.T) {
	slice := make([]foo, 100)

	for i := len(slice) - 1; i >= 0; i-- {
		slice[i].f1 = i
	}

	q := Sort(func(i, j foo) bool {
		return i.f1 < j.f1
	})(FromSlice(slice))

	j := 0
	next := q.Iterate()
	for item, ok := next(); ok; item, ok = next() {
		if item.f1 != j {
			t.Errorf("Sort()[%v]=%v expected %v", j, item, foo{f1: j})
		}

		j++
	}
}
