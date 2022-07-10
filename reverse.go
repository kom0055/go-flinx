package flinx

// Reverse inverts the order of the elements in a collection.
//
// Unlike OrderBy, this sorting method does not consider the actual values
// themselves in determining the order. Rather, it just returns the elements in
// the reverse order from which they are produced by the underlying source.
func Reverse[T any](q Query[T]) Query[T] {
	return Query[T]{
		Iterate: func() Iterator[T] {
			next := q.Iterate()

			items := []T{}
			for item, ok := next(); ok; item, ok = next() {
				items = append(items, item)
			}

			index := len(items) - 1
			return func() (item T, ok bool) {
				if index < 0 {
					return
				}

				item, ok = items[index], true
				index--
				return
			}
		},
	}

}
