package flinx

import "github.com/kom0055/go-flinx/hashset"

// Except produces the set difference of two sequences. The set difference is
// the members of the first sequence that don't appear in the second sequence.
func Except[T comparable](q, q2 Query[T]) Query[T] {
	return Query[T]{
		Iterate: func() Iterator[T] {
			next := q.Iterate()

			next2 := q2.Iterate()
			set := hashset.NewAny[T]()
			for i, ok := next2(); ok; i, ok = next2() {
				set.Insert(i)
			}

			return func() (item T, ok bool) {
				for item, ok = next(); ok; item, ok = next() {
					if !set.Has(item) {
						return
					}
				}

				return
			}
		},
	}

}

// ExceptBy invokes a transform function on each element of a collection and
// produces the set difference of two sequences. The set difference is the
// members of the first sequence that don't appear in the second sequence.
func ExceptBy[T any, V comparable](selector func(T) V) func(q, q2 Query[T]) Query[T] {
	return func(q, q2 Query[T]) Query[T] {
		return Query[T]{
			Iterate: func() Iterator[T] {
				next := q.Iterate()

				next2 := q2.Iterate()
				set := hashset.NewAny[V]()
				for i, ok := next2(); ok; i, ok = next2() {
					s := selector(i)
					set.Insert(s)

				}

				return func() (item T, ok bool) {
					for item, ok = next(); ok; item, ok = next() {
						s := selector(item)
						if !set.Has(s) {
							return
						}

					}

					return
				}
			},
		}
	}

}
