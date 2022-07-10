package flinx

import "github.com/kom0055/flinx/hashset"

// Distinct method returns distinct elements from a collection. The result is an
// unordered collection that contains no duplicate values.
func Distinct[T comparable](q Query[T]) Query[T] {
	return Query[T]{
		Iterate: func() Iterator[T] {
			next := q.Iterate()
			set := hashset.NewAny[T]()

			return func() (item T, ok bool) {
				for item, ok = next(); ok; item, ok = next() {
					if !set.Has(item) {
						set.Insert(item)
						return
					}
				}

				return
			}
		},
	}

}

// DistinctOrderedQuery method returns distinct elements from a collection. The result is an
// ordered collection that contains no duplicate values.
//
// NOTE: DistinctOrderedQuery method on OrderedQuery type has better performance than
// Distinct method on Query type.
func DistinctOrderedQuery[T comparable](oq OrderedQuery[T]) OrderedQuery[T] {
	return OrderedQuery[T]{
		Query: Query[T]{
			Iterate: func() Iterator[T] {
				next := oq.Iterate()
				var prev T

				return func() (item T, ok bool) {
					for item, ok = next(); ok; item, ok = next() {
						if item != prev {
							prev = item
							return
						}
					}

					return
				}
			},
		},
	}

}

// DistinctBy method returns distinct elements from a collection. This method
// executes selector function for each element to determine a value to compare.
// The result is an unordered collection that contains no duplicate values.
func DistinctBy[T any, V comparable](selector func(T) V) func(q Query[T]) Query[T] {
	return func(q Query[T]) Query[T] {
		return Query[T]{
			Iterate: func() Iterator[T] {
				next := q.Iterate()
				set := hashset.NewAny[V]()

				return func() (item T, ok bool) {
					for item, ok = next(); ok; item, ok = next() {
						s := selector(item)
						if !set.Has(s) {
							set.Insert(s)
							return
						}
					}

					return
				}
			},
		}
	}

}
