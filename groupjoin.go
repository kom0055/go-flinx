package flinx

// GroupJoin correlates the elements of two collections based on key equality,
// and groups the results.
//
// This method produces hierarchical results, which means that elements from
// outer query are paired with collections of matching elements from inner.
// GroupJoin enables you to base your results on a whole set of matches for each
// element of outer query.
//
// The resultSelector function is called only one time for each outer element
// together with a collection of all the inner elements that match the outer
// element. This differs from the Join method, in which the result selector
// function is invoked on pairs that contain one element from outer and one
// element from inner.
//
// GroupJoin preserves the order of the elements of outer, and for each element
// of outer, the order of the matching elements from inner.
func GroupJoin[T, V, O any, K comparable](
	outerKeySelector func(T) K,
	innerKeySelector func(V) K,
	resultSelector func(outer T, inners []V) O) func(q Query[T], inner Query[V]) Query[O] {
	return func(q Query[T], inner Query[V]) Query[O] {
		return Query[O]{
			Iterate: func() Iterator[O] {
				outernext := q.Iterate()
				innernext := inner.Iterate()

				innerLookup := make(map[K][]V)
				for innerItem, ok := innernext(); ok; innerItem, ok = innernext() {
					innerKey := innerKeySelector(innerItem)
					innerLookup[innerKey] = append(innerLookup[innerKey], innerItem)
				}

				return func() (item O, ok bool) {
					var tItem T
					if tItem, ok = outernext(); !ok {
						return
					}

					if group, has := innerLookup[outerKeySelector(tItem)]; !has {
						item = resultSelector(tItem, []V{})
					} else {
						item = resultSelector(tItem, group)
					}

					return
				}
			},
		}
	}

}
