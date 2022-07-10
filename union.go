package flinx

import "github.com/kom0055/flinx/hashset"

// Union produces the set union of two collections.
//
// This method excludes duplicates from the return set. This is different
// behavior to the Concat method, which returns all the elements in the input
// collection including duplicates.
func Union[T comparable](q, q2 Query[T]) Query[T] {
	return Query[T]{
		Iterate: func() Iterator[T] {
			next := q.Iterate()
			next2 := q2.Iterate()

			set := hashset.NewAny[T]()
			use1 := true

			return func() (item T, ok bool) {
				if use1 {
					for item, ok = next(); ok; item, ok = next() {
						if !set.Has(item) {
							set.Insert(item)
							return
						}

					}

					use1 = false
				}

				for item, ok = next2(); ok; item, ok = next2() {
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
