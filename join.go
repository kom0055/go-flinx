package flinx

// Join correlates the elements of two collection based on matching keys.
//
// A join refers to the operation of correlating the elements of two sources of
// information based on a common key. Join brings the two information sources
// and the keys by which they are matched together in one method call. This
// differs from the use of SelectMany, which requires more than one method call
// to perform the same operation.
//
// Join preserves the order of the elements of outer collection, and for each of
// these elements, the order of the matching elements of inner.
func Join[T, V, Q any, K comparable](
	outerKeySelector func(T) K,
	innerKeySelector func(V) K,
	resultSelector func(outer T, inner V) Q) func(q Query[T], inner Query[V]) Query[Q] {
	return func(q Query[T], inner Query[V]) Query[Q] {
		return Query[Q]{
			Iterate: func() Iterator[Q] {
				outernext := q.Iterate()
				innernext := inner.Iterate()

				innerLookup := make(map[K][]V)
				for innerItem, ok := innernext(); ok; innerItem, ok = innernext() {
					innerKey := innerKeySelector(innerItem)
					innerLookup[innerKey] = append(innerLookup[innerKey], innerItem)
				}

				var outerItem T
				var innerGroup []V
				innerLen, innerIndex := 0, 0

				return func() (item Q, ok bool) {
					if innerIndex >= innerLen {
						has := false
						for !has {
							outerItem, ok = outernext()
							if !ok {
								return
							}

							innerGroup, has = innerLookup[outerKeySelector(outerItem)]
							innerLen = len(innerGroup)
							innerIndex = 0
						}
					}

					item = resultSelector(outerItem, innerGroup[innerIndex])
					innerIndex++
					return item, true
				}
			},
		}
	}

}
