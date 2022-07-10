package flinx

import "github.com/kom0055/go-flinx/hashset"

// Intersect produces the set intersection of the source collection and the
// provided input collection. The intersection of two sets A and B is defined as
// the set that contains all the elements of A that also appear in B, but no
// other elements.
func Intersect[T comparable](q, q2 Query[T]) Query[T] {
	return Query[T]{
		Iterate: func() Iterator[T] {
			next := q.Iterate()
			next2 := q2.Iterate()

			set := hashset.NewAny[T]()
			for item, ok := next2(); ok; item, ok = next2() {
				set.Insert(item)
			}

			return func() (item T, ok bool) {
				for item, ok = next(); ok; item, ok = next() {
					if set.Has(item) {
						set.Delete(item)
						return
					}

				}

				return
			}
		},
	}

}

// IntersectBy produces the set intersection of the source collection and the
// provided input collection. The intersection of two sets A and B is defined as
// the set that contains all the elements of A that also appear in B, but no
// other elements.
//
// IntersectBy invokes a transform function on each element of both collections.
func IntersectBy[T any, V comparable](selector func(T) V) func(q, q2 Query[T]) Query[T] {
	return func(q, q2 Query[T]) Query[T] {
		return Query[T]{
			Iterate: func() Iterator[T] {
				next := q.Iterate()
				next2 := q2.Iterate()

				set := hashset.NewAny[V]()
				for item, ok := next2(); ok; item, ok = next2() {
					s := selector(item)
					set.Insert(s)
				}

				return func() (item T, ok bool) {

					for item, ok = next(); ok; item, ok = next() {
						s := selector(item)
						if set.Has(s) {
							set.Delete(s)
							return
						}
					}

					return
				}
			},
		}
	}

}
