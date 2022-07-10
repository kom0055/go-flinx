package flinx

import (
	"sort"
)

type order[T any, V any] struct {
	selector func(T) V
	desc     bool
}

// OrderedQuery is the type returned from OrderBy, OrderByDescending ThenBy and
// ThenByDescending functions.
type OrderedQuery[T any] struct {
	Query[T]
	original Query[T]
	cmp      func(T, T) int
}

// OrderBy sorts the elements of a collection in ascending order. Elements are
// sorted according to a key.
func OrderBy[T, V any](compare func(V, V) int, selector func(T) V) func(q Query[T]) OrderedQuery[T] {

	return func(q Query[T]) OrderedQuery[T] {
		return ThenBy[T, V](compare, selector)(OrderedQuery[T]{
			Query:    q,
			original: q,
			cmp:      nil,
		})
	}

}

// OrderByDescending sorts the elements of a collection in descending order.
// Elements are sorted according to a key.
func OrderByDescending[T, V any](compare func(V, V) int, selector func(T) V) func(q Query[T]) OrderedQuery[T] {
	cmp := func(v1, v2 V) int {
		return -compare(v1, v2)
	}
	return OrderBy[T, V](cmp, selector)

}

// ThenBy performs a subsequent ordering of the elements in a collection in
// ascending order. This method enables you to specify multiple sort criteria by
// applying any number of ThenBy or ThenByDescending methods.
func ThenBy[T, V any](compare func(V, V) int, selector func(T) V) func(oq OrderedQuery[T]) OrderedQuery[T] {

	return func(oq OrderedQuery[T]) OrderedQuery[T] {
		cmp := func(t1, t2 T) int {
			if oq.cmp != nil {
				if c := oq.cmp(t1, t2); c != 0 {
					return c
				}
			}
			if compare != nil {
				v1, v2 := selector(t1), selector(t2)
				return compare(v1, v2)
			}

			return 0
		}
		return OrderedQuery[T]{
			original: oq.original,
			Query: Query[T]{
				Iterate: func() Iterator[T] {
					items := sortQuery[T](cmp)(oq.original)
					length := len(items)
					index := 0

					return func() (item T, ok bool) {
						ok = index < length
						if ok {
							item = items[index]
							index++
						}

						return
					}
				},
			},
			cmp: cmp,
		}
	}

}

// ThenByDescending performs a subsequent ordering of the elements in a
// collection in descending order. This method enables you to specify multiple
// sort criteria by applying any number of ThenBy or ThenByDescending methods.
func ThenByDescending[T, V any](compare func(V, V) int, selector func(T) V) func(oq OrderedQuery[T]) OrderedQuery[T] {
	cmp := func(v1, v2 V) int {
		return -compare(v1, v2)
	}
	return ThenBy[T, V](cmp, selector)

}

// Sort returns a new query by sorting elements with provided less function in
// ascending order. The comparer function should return true if the parameter i
// is less than j. While this method is uglier than chaining OrderBy,
// OrderByDescending, ThenBy and ThenByDescending methods, it's performance is
// much better.
func Sort[T any](less func(i, j T) bool) func(q Query[T]) Query[T] {
	lessFn := lessSort[T](less)
	return func(q Query[T]) Query[T] {
		return Query[T]{
			Iterate: func() Iterator[T] {
				items := lessFn(q)
				length := len(items)
				index := 0

				return func() (item T, ok bool) {
					ok = index < length
					if ok {
						item = items[index]
						index++
					}

					return
				}
			},
		}
	}

}

type sorter[T any] struct {
	items []T
	less  func(i, j T) bool
}

func (s sorter[T]) Len() int {
	return len(s.items)
}

func (s sorter[T]) Swap(i, j int) {
	s.items[i], s.items[j] = s.items[j], s.items[i]
}

func (s sorter[T]) Less(i, j int) bool {
	return s.less(s.items[i], s.items[j])
}

func sortQuery[T any](cmp func(T, T) int) func(q Query[T]) (r []T) {
	return func(q Query[T]) (r []T) {
		next := q.Iterate()
		for item, ok := next(); ok; item, ok = next() {
			r = append(r, item)
		}

		if len(r) == 0 {
			return
		}

		s := sorter[T]{
			items: r,
			less: func(i, j T) bool {
				return cmp(i, j) < 0
			},
		}

		sort.Sort(s)
		return
	}

}

func lessSort[T any](less func(i, j T) bool) func(q Query[T]) (r []T) {
	return func(q Query[T]) (r []T) {
		next := q.Iterate()
		for item, ok := next(); ok; item, ok = next() {
			r = append(r, item)
		}

		s := sorter[T]{items: r, less: less}

		sort.Sort(s)
		return
	}

}
