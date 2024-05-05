package flinx

func Where[T any](predicates ...func(T) bool) func(q Query[T]) Query[T] {
	//predicateIdx := func(_ int, item T) bool {
	//	return predicate(item)
	//}
	//return WhereIndexed(predicateIdx)
	predicate := Predicates(predicates...)
	return func(q Query[T]) Query[T] {
		return Query[T]{
			Iterate: func() Iterator[T] {
				next := q.Iterate()
				return func() (item T, ok bool) {
					for item, ok = next(); ok; item, ok = next() {
						if predicate(item) {
							return
						}
					}
					return
				}
			},
		}
	}

}

// WhereIndexed filters a collection of values based on a predicate. Each
// element's index is used in the logic of the predicate function.
//
// The first argument represents the zero-based index of the element within
// collection. The second argument of predicate represents the element to test.
func WhereIndexed[T any](predicates ...func(int, T) bool) func(q Query[T]) Query[T] {
	predicate := PredicatesIndexed(predicates...)
	return func(q Query[T]) Query[T] {
		return Query[T]{
			Iterate: func() Iterator[T] {
				next := q.Iterate()
				index := 0

				return func() (item T, ok bool) {
					for item, ok = next(); ok; item, ok = next() {
						if predicate(index, item) {
							index++
							return
						}
						index++
					}
					return
				}
			},
		}
	}

}
