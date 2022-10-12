package flinx

// Skip bypasses a specified number of elements in a collection and then returns
// the remaining elements.
func Skip[T any](q Query[T], count int) Query[T] {

	return Query[T]{
		Iterate: func() Iterator[T] {
			next := q.Iterate()
			n := count

			return func() (item T, ok bool) {
				for ; n > 0; n-- {
					item, ok = next()
					if !ok {
						return
					}
				}

				return next()
			}
		},
	}

}

// SkipWhile bypasses elements in a collection as long as a specified condition
// is true and then returns the remaining elements.
//
// This method tests each element by using predicate and skips the element if
// the result is true. After the predicate function returns false for an
// element, that element and the remaining elements in source are returned and
// there are no more invocations of predicate.
func SkipWhile[T any](predicate func(T) bool) func(q Query[T]) Query[T] {

	return func(q Query[T]) Query[T] {
		return Query[T]{
			Iterate: func() Iterator[T] {
				next := q.Iterate()
				ready := false

				return func() (item T, ok bool) {
					for !ready {
						item, ok = next()
						if !ok {
							return
						}

						ready = !predicate(item)
						if ready {
							return
						}
					}

					return next()
				}
			},
		}
	}

}

// SkipWhileIndexed bypasses elements in a collection as long as a specified
// condition is true and then returns the remaining elements. The element's
// index is used in the logic of the predicate function.
//
// This method tests each element by using predicate and skips the element if
// the result is true. After the predicate function returns false for an
// element, that element and the remaining elements in source are returned and
// there are no more invocations of predicate.
func SkipWhileIndexed[T any](predicate func(int, T) bool) func(q Query[T]) Query[T] {
	return func(q Query[T]) Query[T] {
		return Query[T]{
			Iterate: func() Iterator[T] {
				next := q.Iterate()
				ready := false
				index := 0

				return func() (item T, ok bool) {
					for !ready {
						item, ok = next()
						if !ok {
							return
						}

						ready = !predicate(index, item)
						if ready {
							return
						}

						index++
					}

					return next()
				}
			},
		}
	}

}
