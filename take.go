package flinx

// Take returns a specified number of contiguous elements from the start of a
// collection.
func Take[T any](q Query[T], count int) Query[T] {

	return Query[T]{
		Iterate: func() Iterator[T] {
			next := q.Iterate()
			n := count

			return func() (item T, ok bool) {
				if n <= 0 {
					return
				}

				n--
				return next()
			}
		},
	}

}

// TakeWhile returns elements from a collection as long as a specified condition
// is true, and then skips the remaining elements.
func TakeWhile[T any](predicates ...func(T) bool) func(q Query[T]) Query[T] {
	predicate := Predicates(predicates...)
	return func(q Query[T]) Query[T] {
		return Query[T]{
			Iterate: func() Iterator[T] {
				next := q.Iterate()
				done := false

				return func() (item T, ok bool) {
					if done {
						return
					}

					item, ok = next()
					if !ok {
						done = true
						return
					}

					if predicate(item) {
						return
					}

					done = true
					var r T
					return r, false
				}
			},
		}
	}

}

// TakeWhileIndexed returns elements from a collection as long as a specified
// condition is true. The element's index is used in the logic of the predicate
// function. The first argument of predicate represents the zero-based index of
// the element within collection. The second argument represents the element to
// test.
func TakeWhileIndexed[T any](predicates ...func(int, T) bool) func(q Query[T]) Query[T] {
	predicate := PredicatesIndexed(predicates...)
	return func(q Query[T]) Query[T] {
		return Query[T]{
			Iterate: func() Iterator[T] {
				next := q.Iterate()
				done := false
				index := 0

				return func() (item T, ok bool) {
					if done {
						return
					}

					item, ok = next()
					if !ok {
						done = true
						return
					}

					if predicate(index, item) {
						index++
						return
					}

					done = true
					var r T
					return r, false
				}
			},
		}
	}

}
