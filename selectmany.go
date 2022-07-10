package flinx

// SelectMany projects each element of a collection to a Query, iterates and
// flattens the resulting collection into one collection.
func SelectMany[T, V any](selector func(T) Query[V]) func(q Query[T]) Query[V] {

	selectorIdx := func(_ int, t T) Query[V] {
		return selector(t)
	}

	return SelectManyIndexed[T, V](selectorIdx)

	//return func(q Query[T]) Query[V] {
	//	return Query[V]{
	//		Iterate: func() Iterator[V] {
	//			outernext := q.Iterate()
	//			var inner *T
	//			var innernext Iterator[V]
	//			return func() (item V, ok bool) {
	//				for !ok {
	//					if inner == nil {
	//						inner, ok = func() (*T, bool) {
	//							itemT, exist := outernext()
	//							return &itemT, exist
	//						}()
	//						if !ok {
	//							return
	//						}
	//
	//						innernext = selector(*inner).Iterate()
	//					}
	//
	//					item, ok = innernext()
	//					if !ok {
	//						inner = nil
	//					}
	//				}
	//
	//				return
	//			}
	//		},
	//	}
	//}
}

// SelectManyIndexed projects each element of a collection to a Query, iterates
// and flattens the resulting collection into one collection.
//
// The first argument to selector represents the zero-based index of that
// element in the source collection. This can be useful if the elements are in a
// known order and you want to do something with an element at a particular
// index, for example. It can also be useful if you want to retrieve the index
// of one or more elements. The second argument to selector represents the
// element to process.
func SelectManyIndexed[T, V any](selector func(int, T) Query[V]) func(q Query[T]) Query[V] {

	return func(q Query[T]) Query[V] {
		return Query[V]{
			Iterate: func() Iterator[V] {
				outernext := q.Iterate()
				index := 0
				var inner *T
				var innernext Iterator[V]

				return func() (item V, ok bool) {
					for !ok {
						if inner == nil {
							inner, ok = func() (*T, bool) {
								itemT, exist := outernext()
								return &itemT, exist
							}()
							if !ok {
								return
							}

							innernext = selector(index, *inner).Iterate()
							index++
						}

						item, ok = innernext()
						if !ok {
							inner = nil
						}
					}

					return
				}
			},
		}
	}

}

// SelectManyBy projects each element of a collection to a Query, iterates and
// flattens the resulting collection into one collection, and invokes a result
// selector function on each element therein.
func SelectManyBy[T, V, O any](selector func(T) Query[V],
	resultSelector func(V, T) O) func(q Query[T]) Query[O] {

	return func(q Query[T]) Query[O] {
		return Query[O]{
			Iterate: func() Iterator[O] {
				outernext := q.Iterate()
				var outer *T
				var innernext Iterator[V]

				return func() (item O, ok bool) {
					var v V
					for !ok {
						if outer == nil {
							outer, ok = func() (*T, bool) {
								outerTmp, exist := outernext()
								return &outerTmp, exist
							}()
							if !ok {
								return
							}

							innernext = selector(*outer).Iterate()
						}

						v, ok = innernext()
						if !ok {
							outer = nil
						}
					}

					item = resultSelector(v, *outer)
					return
				}
			},
		}
	}

}

// SelectManyByIndexed projects each element of a collection to a Query,
// iterates and flattens the resulting collection into one collection, and
// invokes a result selector function on each element therein. The index of each
// source element is used in the intermediate projected form of that element.

func SelectManyByIndexed[T, V, O any](selector func(int, T) Query[V],
	resultSelector func(V, T) O) func(q Query[T]) Query[O] {
	return func(q Query[T]) Query[O] {
		return Query[O]{
			Iterate: func() Iterator[O] {
				outernext := q.Iterate()
				index := 0
				var outer *T
				var innernext Iterator[V]

				return func() (item O, ok bool) {
					var v V
					for !ok {
						if outer == nil {
							outer, ok = func() (*T, bool) {
								outerTmp, exist := outernext()
								return &outerTmp, exist
							}()
							if !ok {
								return
							}

							innernext = selector(index, *outer).Iterate()
							index++
						}

						v, ok = innernext()
						if !ok {
							outer = nil
						}
					}

					item = resultSelector(v, *outer)
					return
				}
			},
		}
	}

}
