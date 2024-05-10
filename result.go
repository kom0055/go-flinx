package flinx

import (
	"math"

	"golang.org/x/exp/constraints"
	_ "golang.org/x/exp/constraints"
)

// All determines whether all elements of a collection satisfy a condition.
func All[T any](predicates ...func(T) bool) func(q Query[T]) bool {
	predicate := Predicates(predicates...)
	return func(q Query[T]) bool {
		next := q.Iterate()

		for item, ok := next(); ok; item, ok = next() {
			if !predicate(item) {
				return false
			}
		}

		return true
	}

}

// Any determines whether any element of a collection exists.
func Any[T any](q Query[T]) bool {

	_, ok := q.Iterate()()
	return ok

}

// AnyWith determines whether any element of a collection satisfies a condition.
func AnyWith[T any](predicates ...func(T) bool) func(q Query[T]) bool {
	predicate := Predicates(predicates...)
	return func(q Query[T]) bool {
		next := q.Iterate()

		for item, ok := next(); ok; item, ok = next() {
			if predicate(item) {
				return true
			}
		}

		return false
	}

}

// Average computes the average of a collection of numeric values.
func Average[T constraints.Integer | constraints.Float](q Query[T]) (r float64) {
	next := q.Iterate()
	item, ok := next()
	if !ok {
		return math.NaN()
	}

	sum := item
	n := 1
	for item, ok = next(); ok; item, ok = next() {
		sum += item
		n++
	}
	r = float64(sum)
	return r / float64(n)

}

// Contains determines whether a collection contains a specified element.
func Contains[T comparable](value T) func(q Query[T]) bool {
	return func(q Query[T]) bool {
		next := q.Iterate()

		for item, ok := next(); ok; item, ok = next() {
			if item == value {
				return true
			}
		}

		return false
	}

}

// Count returns the number of elements in a collection.
func Count[T any](q Query[T]) (r int) {
	next := q.Iterate()

	for _, ok := next(); ok; _, ok = next() {
		r++
	}
	return
}

// CountWith returns a number that represents how many elements in the specified
// collection satisfy a condition.
func CountWith[T any](predicates ...func(T) bool) func(q Query[T]) (r int) {
	predicate := Predicates(predicates...)
	return func(q Query[T]) (r int) {
		next := q.Iterate()

		for item, ok := next(); ok; item, ok = next() {
			if predicate(item) {
				r++
			}
		}

		return
	}

}

// First returns the first element of a collection.
func First[T any](q Query[T]) (T, bool) {
	return q.Iterate()()
}

// FirstWith returns the first element of a collection that satisfies a
// specified condition.
func FirstWith[T any](predicates ...func(T) bool) func(q Query[T]) (T, bool) {
	predicate := Predicates(predicates...)
	return func(q Query[T]) (T, bool) {
		next := q.Iterate()

		for item, ok := next(); ok; item, ok = next() {
			if predicate(item) {
				return item, ok
			}
		}
		var r T

		return r, false
	}

}

// ForEach performs the specified action on each element of a collection.
func ForEach[T any](action func(T)) func(q Query[T]) {
	return func(q Query[T]) {
		next := q.Iterate()

		for item, ok := next(); ok; item, ok = next() {
			action(item)
		}
	}

}

// ForEachIndexed performs the specified action on each element of a collection.
//
// The first argument to action represents the zero-based index of that
// element in the source collection. This can be useful if the elements are in a
// known order and you want to do something with an element at a particular
// index, for example. It can also be useful if you want to retrieve the index
// of one or more elements. The second argument to action represents the
// element to process.
func ForEachIndexed[T any](action func(int, T)) func(q Query[T]) {
	return func(q Query[T]) {
		next := q.Iterate()
		index := 0

		for item, ok := next(); ok; item, ok = next() {
			action(index, item)
			index++
		}
	}

}

// Last returns the last element of a collection.
func Last[T any](q Query[T]) (r T, exist bool) {
	next := q.Iterate()

	for item, ok := next(); ok; item, ok = next() {
		r = item
		exist = true
	}

	return
}

// LastWith returns the last element of a collection that satisfies a specified
// condition.
func LastWith[T any](predicates ...func(T) bool) func(q Query[T]) (r T, exist bool) {
	predicate := Predicates(predicates...)
	return func(q Query[T]) (r T, exist bool) {
		next := q.Iterate()

		for item, ok := next(); ok; item, ok = next() {
			if predicate(item) {
				r = item
				exist = true
			}
		}

		return
	}

}

// Max returns the maximum value in a collection of values.
func Max[T any](compare func(t1, t2 T) int) func(q Query[T]) (r T, exist bool) {
	return func(q Query[T]) (r T, exist bool) {
		next := q.Iterate()
		item, ok := next()
		if !ok {
			return
		}

		r = item
		exist = true

		for item, ok := next(); ok; item, ok = next() {
			if compare(item, r) > 0 {
				r = item
				exist = true
			}
		}

		return
	}

}

// Min returns the minimum value in a collection of values.
func Min[T any](compare func(t1, t2 T) int) func(q Query[T]) (r T, exist bool) {
	return func(q Query[T]) (r T, exist bool) {
		next := q.Iterate()
		item, ok := next()
		if !ok {
			return
		}

		r = item
		exist = true

		for item, ok := next(); ok; item, ok = next() {
			if compare(item, r) < 0 {
				r = item
				exist = true
			}
		}

		return
	}

}

// Results iterates over a collection and returnes slice of interfaces
func Results[T any](q Query[T]) (r []T) {
	next := q.Iterate()

	for item, ok := next(); ok; item, ok = next() {
		r = append(r, item)
	}

	return
}

// SequenceEqual determines whether two collections are equal.
func SequenceEqual[T comparable](q, q2 Query[T]) bool {
	next := q.Iterate()
	next2 := q2.Iterate()

	for item, ok := next(); ok; item, ok = next() {
		item2, ok2 := next2()
		if !ok2 || item != item2 {
			return false
		}
	}

	_, ok2 := next2()
	return !ok2

}

// Single returns the only element of a collection, and nil if there is not
// exactly one element in the collection.
func Single[T any](q Query[T]) (r T, found bool) {
	next := q.Iterate()
	item, ok := next()
	if !ok {
		return
	}

	_, ok = next()
	if ok {
		return
	}

	return item, true
}

// SingleWith returns the only element of a collection that satisfies a
// specified condition, and nil if more than one such element exists.
func SingleWith[T any](predicates ...func(T) bool) func(q Query[T]) (r T, found bool) {
	predicate := Predicates(predicates...)
	return func(q Query[T]) (r T, found bool) {
		next := q.Iterate()

		for item, ok := next(); ok; item, ok = next() {
			if predicate(item) {
				if found {
					var v T
					return v, false
				}
				found = true
				r = item

			}
		}

		return
	}

}

// Sum computes the sum of a collection of numeric values.
//
// Values can be of any integer type: int, int8, int16, int32, int64. The result
// is int64. Method returns zero if collection contains no elements.
func Sum[T constraints.Integer | constraints.Float](q Query[T]) (r T) {
	next := q.Iterate()
	item, ok := next()
	if !ok {
		return 0
	}
	r = item
	for item, ok = next(); ok; item, ok = next() {
		r += item
	}

	return
}

// ToChannel iterates over a collection and outputs each element to a channel,
// then closes it.
func ToChannel[T any](q Query[T], result chan<- T) {
	defer close(result)
	next := q.Iterate()
	for item, ok := next(); ok; item, ok = next() {
		result <- item
	}

}

// ToMap iterates over a collection and populates result map with elements.
// Collection elements have to be of KeyValue type to use this method. To
// populate a map with elements of different type use ToMapBy method. ToMap
// doesn't empty the result map before populating it.
func ToMap[K comparable, V any](q Query[KeyValue[K, V]]) map[K]V {
	m := map[K]V{}

	next := q.Iterate()

	for item, ok := next(); ok; item, ok = next() {

		m[item.Key] = item.Value
	}
	return m
}

func ToMapFromGroup[K comparable, V any](q Query[Group[K, V]]) map[K][]V {
	m := map[K][]V{}

	next := q.Iterate()

	for item, ok := next(); ok; item, ok = next() {
		m[item.Key] = append(m[item.Key], item.Group...)
	}
	return m
}

// ToMapBy iterates over a collection and populates the result map with
// elements. Functions keySelector and valueSelector are executed for each
// element of the collection to generate key and value for the map. Generated
// key and value types must be assignable to the map's key and value types.
// ToMapBy doesn't empty the result map before populating it.
func ToMapBy[K comparable, V, T any](keySelector func(T) K, valueSelector func(T) V) func(q Query[T]) map[K]V {
	return func(q Query[T]) map[K]V {
		m := map[K]V{}

		next := q.Iterate()

		for item, ok := next(); ok; item, ok = next() {

			m[keySelector(item)] = valueSelector(item)
		}
		return m
	}

}

// ToSlice iterates over a collection and saves the results in the slice pointed
// by v. It overwrites the existing slice, starting from index 0.
//
// If the slice pointed by v has sufficient capacity, v will be pointed to a
// resliced slice. If it does not, a new underlying array will be allocated and
// v will point to it.
func ToSlice[T any](q Query[T]) []T {
	r := []T{}
	next := q.Iterate()
	for item, ok := next(); ok; item, ok = next() {
		r = append(r, item)
	}
	return r

}

func ToString(q Query[rune]) string {
	return string(ToSlice[rune](q))
}
