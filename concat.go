package flinx

// Append inserts an item to the end of a collection, so it becomes the last
// item.
func Append[T any](items ...T) func(q Query[T]) Query[T] {
	return func(q Query[T]) Query[T] {
		length := len(items)
		var defaultValue T
		return Query[T]{
			Iterate: func() Iterator[T] {
				next := q.Iterate()
				index := 0

				return func() (T, bool) {
					i, ok := next()
					if ok {
						return i, ok
					}
					if index < length {
						idx := index
						index++

						return items[idx], true
					}

					return defaultValue, false
				}
			},
		}
	}

}

// Concat concatenates two collections.
//
// The Concat method differs from the Union method because the Concat method
// returns all the original elements in the input sequences. The Union method
// returns only unique elements.
func Concat[T any](q, q2 Query[T]) Query[T] {
	return Query[T]{
		Iterate: func() Iterator[T] {
			next := q.Iterate()
			next2 := q2.Iterate()
			use1 := true

			return func() (item T, ok bool) {
				if use1 {
					item, ok = next()
					if ok {
						return
					}

					use1 = false
				}

				return next2()
			}
		},
	}

}

// Prepend inserts an item to the beginning of a collection, so it becomes the
// first item.
func Prepend[T any](items ...T) func(q Query[T]) Query[T] {
	return func(q Query[T]) Query[T] {
		length := len(items)
		return Query[T]{
			Iterate: func() Iterator[T] {
				next := q.Iterate()
				index := 0

				return func() (T, bool) {
					if index < length {
						idx := index
						index++
						return items[idx], true

					}

					return next()
				}
			},
		}
	}

}
